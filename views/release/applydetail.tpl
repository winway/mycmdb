    <div class="modal-header">
      <button type="button" class="close" data-dismiss="modal" aria-hidden="true">
        &times;
      </button>
      <h4 class="modal-title" id="myModalLabel">
        流程处理-{{ .apply.Subject }}<br/ >
        当前步骤-{{ .currentStep.Description }}
      </h4>
    </div>
    <form id="app" class="form-horizontal" enctype="multipart/form-data" method="post" action="">
      <div class="modal-body">
        <input v-model="applyId" type="hidden" name="applyId">
        <input v-model="currentStep" type="hidden" name="currentStep">

        {{ if eq .currentStep.Owner .email }}
        <div class="form-group">
          <label for="Content" class="col-sm-3 control-label">说明（可选）：</label>
          <div class="col-sm-8 row-sm-8">
            <textarea v-model="Content" class="form-control" style="height:300px" name="Content" placeholder="请填写相关信息"></textarea>
          </div>
        </div>

        <div class="form-group">
          <label for="Content" class="col-sm-3 control-label">附件（可选）：</label>
          <div class="col-sm-8">
            <input v-model="myfile" id="myfile" name="myfile" type="file" />
          </div>
        </div>

        {{ if eq .currentStep.IsLast 1}}
        <div class="form-group">
          <input v-model="nextOwner" type="hidden" name="nextOwner">
          {{else}}
          <label  class="col-sm-3 control-label" for="nextOwner">转交下一负责人：</label>
          <div class="col-sm-8">
            <select v-model="nextOwner" name="nextOwner" class="form-control" title="请选择转交人" data-style="btn-default">
              {{range $k, $v := .emails}}
              <option value={{ $v.Email }}>{{ $v.Email }}</option>
              {{end}}
            </select>
          </div>
        </div>
        {{ end }}
        {{ end }}

        <br />
        <h3>进度详情</h3>
        <table class="table">
          <tr>
            <th>序号</th>
            <th>描述</th>
            <th>负责人</th>
            <th>状态</th>
            <th>完成时间</th>
            <th>查看详情</th>
          </tr>
          <tr v-for="(step, index) in Steps">
            <td :class="(step.Status == '已完成')? 'btn-success':'btn-warning'" v-text="step.StepId"></td>
            <td :class="(step.Status == '已完成')? 'btn-success':'btn-warning'" v-text="step.Description"></td>
            <td :class="(step.Status == '已完成')? 'btn-success':'btn-warning'" v-text="step.Owner"></td>
            <td :class="(step.Status == '已完成')? 'btn-success':'btn-warning'" v-text="step.Status"></td>
            <td :class="(step.Status == '已完成')? 'btn-success':'btn-warning'" v-text="step.OperateTime"></td>
            <td :class="(step.Status == '已完成')? 'btn-success':'btn-warning'"><a v-bind:href="'{{ urlfor "ReleaseController.ApplyViewPage" ":id" "" }}' + step.StepId" target="_blank">查看详情</a></td>
          </tr>
        </table>

      </div>

      <div class="modal-footer">
        <button type="button" class="btn btn-default" data-dismiss="modal">关闭</button>
        {{ if eq .currentStep.Owner .email }}
        <button type="button" v-on:click="confirm" class="btn btn-primary">确认提交</button>
        {{ end }}
      </div>
    </form>

    <script type="text/javascript">
      var vue = new Vue({
        el: '#app',
        data: {
          Apply: {{ .apply }},
          applyId: {{ .apply.Id }},
          Steps: {{ .steps }},
          currentStep: {{ .currentStep.StepId }},
          Content: "",
          myfile: "",
          nextOwner: "0",
        },
        methods: {
          confirm: function() {
            var self = this;
            $.ajax({
              type:"POST",
              url:"/release/detail/{{ .apply.Id }}",
              data:new FormData($('#app')[0]),
              processData: false,
              contentType: false,
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
        }
      });

    </script>