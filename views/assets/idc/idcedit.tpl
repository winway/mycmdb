    <div class="modal-header">
      <button type="button" class="close" data-dismiss="modal" aria-hidden="true">
        &times;
      </button>
      <h4 class="modal-title" id="myModalLabel">
        {{ .title }}
      </h4>
    </div>
    <form class="form-horizontal" method="post" action="{{ urlfor "IdcController.SaveIdc" }}">
      <div class="modal-body">
        <input type="hidden" name="Id" value="{{ .idc.Id }}">

        <div class="form-group">
          <label for="name" class="col-sm-2 control-label">名称：</label>
          <div class="col-sm-10">
            <input type="text" class="form-control" name="Name" placeholder="请输入机房名称" value="{{ .idc.Name }}">
          </div>
        </div>

        <div class="form-group">
          <label for="address" class="col-sm-2 control-label">地址：</label>
          <div class="col-sm-10">
            <input type="text" class="form-control" name="Address" placeholder="请输入机房地址" value="{{ .idc.Address }}">
          </div>
        </div>

        <div class="form-group">
          <label for="operator" class="col-sm-2 control-label">运营商：</label>
          <div class="col-sm-10">
            <input type="text" class="form-control" name="Operator" placeholder="请输入运营商" value="{{ .idc.Operator }}">
          </div>
        </div>

        <div class="form-group">
          <label for="linkman" class="col-sm-2 control-label">联系人：</label>
          <div class="col-sm-10">
            <input type="text" class="form-control" name="Linkman" placeholder="请输入联系人" value="{{ .idc.Linkman }}">
          </div>
        </div>

        <div class="form-group">
          <label for="phone" class="col-sm-2 control-label">联系电话：</label>
          <div class="col-sm-10">
            <input type="text" class="form-control" name="Phone" placeholder="请输入联系电话" value="{{ .idc.Phone }}">
          </div>
        </div>

        <div class="form-group">
          <label for="comment" class="col-sm-2 control-label">备注：</label>
          <div class="col-sm-10">
            <input type="text" class="form-control" name="Comment" placeholder="请输入备注" value="{{ .idc.Comment }}">
          </div>
        </div>
      </div>

      <div class="modal-footer">
        <button type="button" class="btn btn-default" data-dismiss="modal">关闭</button>
        <button type="submit" class="btn btn-primary">保存</button>
      </div>
    </form>