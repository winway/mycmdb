    <div class="modal-header">
      <button type="button" class="close" data-dismiss="modal" aria-hidden="true">
        &times;
      </button>
      <h4 class="modal-title" id="myModalLabel">
        代码发布申请
      </h4>
    </div>
    <form id="app" class="form-horizontal" method="post" action="">
      <div class="modal-body">
        <div class="form-group">
          <label for="Subject" class="col-sm-2 control-label">主题：</label>
          <div class="col-sm-5">
            <input v-model="Subject" type="text" class="form-control" name="Subject" placeholder="">
          </div>
        </div>

        <div class="form-group">
          <label  class="col-sm-2 control-label" for="applyType">类型：</label>
          <div class="col-sm-5">
            <select v-model="applyType" name="applyType" class="form-control" title="请选择操作系统" data-style="btn-default">
              <option value=0>新功能上线</option>
              <option value=1>Bug修复</option>
              <option value=2>功能点优化</option>
            </select>
          </div>
        </div>
      </div>

      <div class="modal-footer">
        <button type="button" class="btn btn-default" data-dismiss="modal">关闭</button>
        <button type="button" v-on:click="confirm" class="btn btn-primary">确认</button>
      </div>
    </form>

    <script type="text/javascript">
      var vue = new Vue({
        el: '#app',
        data: {
          Subject: "",
          applyType: 0
        },
        methods: {
          confirm: function() {
            var self = this;
            $.ajax({
              type:"POST",
              url:"/release/applylist",
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
          }
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