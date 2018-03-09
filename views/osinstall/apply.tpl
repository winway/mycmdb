    <div class="modal-header">
      <button type="button" class="close" data-dismiss="modal" aria-hidden="true">
        &times;
      </button>
      <h4 class="modal-title" id="myModalLabel">
        OS安装申请
      </h4>
    </div>
    <form id="app" class="form-horizontal" method="post" action="">
      <div class="modal-body">
        <div class="form-group">
          <input v-model="Disknum" type="hidden" name="Disknum">

          <label  class="col-sm-2 control-label" for="Sn">SN：</label>
          <div class="col-sm-5">
            <select v-model="Sn" v-on:change='getInfo' name="Sn" class="form-control" title="请选择SN" data-style="btn-default">
              {{range $k, $v := .snSet}}
              <option value={{ $v.Sn }}>{{ $v.Sn }}</option>
              {{end}}
            </select>
          </div>
        </div>

        <div class="form-group">
          <label for="Idc" class="col-sm-2 control-label">机房：</label>
          <div class="col-sm-5">
            <input v-model="Idc" type="text" class="form-control" name="Idc" placeholder="" readonly>
          </div>
        </div>

        <div class="form-group">
          <label for="RemoteCardIp" class="col-sm-2 control-label">远程卡IP：</label>
          <div class="col-sm-5">
            <input v-model="RemoteCardIp" type="text" class="form-control" name="RemoteCardIp" placeholder="请输入数量" readonly>
          </div>
        </div>

        <div class="form-group">
          <label for="Ip" class="col-sm-2 control-label">网卡IP：</label>
          <div class="col-sm-5">
            <textarea v-model="Ip" class="form-control" name="Ip" readonly></textarea>
          </div>
        </div>

        <div class="form-group">
          <label for="Status" class="col-sm-2 control-label">状态：</label>
          <div class="col-sm-5">
            <input v-model="Status" type="text" class="form-control" name="Status" placeholder="请输入数量" readonly>
          </div>
        </div>

        <div class="form-group">
          <label  class="col-sm-2 control-label" for="OsVersion">操作系统：</label>
          <div class="col-sm-5">
            <select v-model="OsVersion" name="OsVersion" class="form-control" title="请选择操作系统" data-style="btn-default">
              <option value=0>CentOS-6.9-x86_64</option>
              <option value=1>centos73-x86_64</option>
            </select>
          </div>
        </div>

        <div class="form-group">
          <label for="Comment" class="col-sm-2 control-label">备注：</label>
          <div class="col-sm-5">
            <textarea v-model="Comment" class="form-control" name="Comment"></textarea>
          </div>
        </div>

        <p>RAID设置</p>
        <hr/>
        <div class="form-inline">
          <div v-for="(disks, index) in Diskmap" v-bind:name="index">
            <div v-for="(disk, key) in disks">
              <input class="form-control" type="checkbox" autocomplete="off" v-bind:name="index + '_' + key" v-bind:value="disk">
              <span class="form-control" v-text="disk"></span>
            </div>
            <select class="form-control" v-bind:name="index">
              <option value ="raid0">raid0</option>
              <option value ="raid1">raid1</option>
              <option value="raid10">raid10</option>
              <option value="raid5">raid5</option>
            </select>
            <input class="form-control" type="button" class="btn btn-default" v-on:click="delRaid(disks)" value="删除"></button>
          </div>
          <input type="button" class="btn btn-default" v-on:click="addRaid" value="添加">
        </div>

      </div>

      <div class="modal-footer">
        <button type="button" class="btn btn-default" data-dismiss="modal">关闭</button>
        <button type="button" v-on:click="confirmInstall" class="btn btn-primary">确认</button>
      </div>
    </form>

    <script type="text/javascript">
      var vue = new Vue({
        el: '#app',
        data: {
          Sn: "",
          Idc: "",
          RemoteCardIp: "",
          Ip: "",
          Status: "",
          OsVersion: 0,
          Comment: "",
          Disknum: 0,
          Disks: {},
          Diskmap: []
        },
        methods: {
          getInfo: function() {
            var self = this;
            $.ajax({
              type:"GET",
              url:"/osinstall/serverinfo",
              data:{
                'Sn': self.Sn
              },
              success: function(data, textStatus, jqXHR) {
                if (data.code != 0) {
                  alert(data.msg);
                } else {
                  self.Idc = data.data.Idc.Name;
                  self.RemoteCardIp = data.data.RemoteCardIp;
                  self.Ip = data.data.Eth1Ip + '\n' + data.data.Eth2Ip + '\n' + data.data.Eth3Ip + '\n' + data.data.Eth4Ip;
                  self.Status = data.data.Status;
                  self.Disks = JSON.parse(data.data.DiskStructure);
                  self.Disknum = getJsonLength(self.Disks);
                }
              }
            });
          },
          confirmInstall: function() {
            var self = this;
            $.ajax({
              type:"POST",
              url:"/osinstall/manifest",
              data:$("#app").serializeArray(),
              success: function(data, textStatus, jqXHR) {
                if (data.code != 0) {
                  alert(data.msg);
                } else {
                  $("#myModal").modal('toggle');
                  Search();
                }
              }
            });
          },
          addRaid: function () {
           this.Diskmap.push(this.Disks);
         },
         delRaid: function (m) {
           this.Diskmap.splice((this.Diskmap.indexOf(m)),1);
         },
       }
     });

      function getJsonLength(jsonData){
        var jsonLength = 0;
        for(var item in jsonData){
          jsonLength++;
        }
        return jsonLength;
      }
    </script>