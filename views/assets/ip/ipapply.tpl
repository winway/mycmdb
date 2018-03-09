    <div class="modal-header">
      <button type="button" class="close" data-dismiss="modal" aria-hidden="true">
        &times;
      </button>
      <h4 class="modal-title" id="myModalLabel">
        IP申请
      </h4>
    </div>
    <form id="app" class="form-horizontal" method="post" action="">
      <div class="modal-body">
        <div class="form-group">
          <label  class="col-sm-2 control-label" for="Idc">机房：</label>
          <div class="col-sm-5">
            <select v-model="idc" name="Idc" class="form-control" title="请选择机房" data-style="btn-default">
              {{range $index, $elem := .names}}
              <option value={{ $elem.Id }}>{{ $elem.Name }}</option>
              {{end}}
            </select>
          </div>
        </div>

        <div class="form-group">
          <label  class="col-sm-2 control-label" for="IpType">类型：</label>
          <div class="col-sm-5">
            <select v-model="ipType" name="IpType" class="form-control" title="请选择类型" data-style="btn-default">
              <option value=0>远程卡IP</option>
              <option value=1>网卡IP</option>
            </select>
          </div>
        </div>

        <div class="form-group">
          <label for="Number" class="col-sm-2 control-label">数量：</label>
          <div class="col-sm-5">
            <input v-model="num" type="text" class="form-control" name="Number" placeholder="请输入数量" value="{{ .idc.Name }}">
          </div>
        </div>

        <div class="form-group">
          <label for="Mail" class="col-sm-2 control-label">申请人邮箱：</label>
          <div class="col-sm-5">
            <input v-model="mail" type="text" class="form-control" name="Mail" placeholder="请输入邮箱" value="{{ .idc.Name }}">
          </div>
        </div>

        <div class="form-group">
          <label for="Ip" class="col-sm-2 control-label">候选IP：</label>
          <div class="col-sm-5">
            <textarea v-model="ips" class="form-control" name="Ip" readonly></textarea>
          </div>
        </div>
      </div>

      <div class="modal-footer">
        <button type="button" class="btn btn-default" data-dismiss="modal">关闭</button>
        <button type="button" v-on:click="getIp" class="btn btn-default">查看IP</button>
        <button type="button" v-on:click="confirmIp" class="btn btn-primary">确认使用</button>
      </div>
    </form>

    <script type="text/javascript">
      var vue = new Vue({
        el: '#app',
        data: {
          idc: 1,
          ipType: 0,
          num: 1,
          mail: "wangwei03@sunlands.com",
          ips: ""
        },
        methods: {
          getIp: function() {
            var self = this;
            $.ajax({
              type:"POST",
              url:"/assets/ip_preapply/",
              data:{
                'idc': self.idc,
                'ipType': self.ipType,
                'num': self.num
              },
              success: function(data, textStatus, jqXHR) {
                if (data.code != 0) {
                  alert(data.msg);
                } else {
                  self.ips = data.data;
                }
              }
            });
          },
          confirmIp: function() {
            var self = this;
            $.ajax({
              type:"POST",
              url:"/assets/ip_apply/",
              data:{
                'ips': self.ips,
                'mail': self.mail
              },
              success: function(data, textStatus, jqXHR) {
                if (data.code != 0) {
                  alert(data.msg);
                } else {
                  $("#myModal").modal('toggle');
                  Search();
                }
              }
            });
          }
        }
      });
    </script>