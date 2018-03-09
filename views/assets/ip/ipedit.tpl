    <div class="modal-header">
      <button type="button" class="close" data-dismiss="modal" aria-hidden="true">
        &times;
      </button>
      <h4 class="modal-title" id="myModalLabel">
        {{ .title }}
      </h4>
    </div>
    <form class="form-horizontal" method="post" action="{{ urlfor "IpController.SaveIp" }}">
      <div class="modal-body">
        <input type="hidden" name="Id" value="{{ .idc.Id }}">

        <div class="form-group">
          <label  class="col-sm-2 control-label" for="Idc">机房：</label>
          <select name="Idc" class="selectpicker" title="请选择机房" data-style="btn-default">
            {{range $index, $elem := .names}}
            {{ if eq 0 $index }}
            <option selected value={{ $elem.Id }}>{{ $elem.Name }}</option>
            {{ else }}
            <option value={{ $elem.Id }}>{{ $elem.Name }}</option>
            {{ end }}
            {{end}}
          </select>
        </div>

        <div class="form-group">
          <label for="Network" class="col-sm-2 control-label">IP/前缀：</label>
          <div class="col-sm-5">
            <input type="text" class="form-control" name="Network" placeholder="请输入IP/前缀" value="{{ .idc.Name }}">
          </div>
        </div>

        <div class="form-group">
          <label for="SubMask" class="col-sm-2 control-label">子网掩码：</label>
          <div class="col-sm-6">
            <input type="text" class="form-control" name="SubMask" placeholder="请输入备注" value="{{ .idc.SubMask }}">
          </div>
        </div>

        <div class="form-group">
          <label for="Gateway" class="col-sm-2 control-label">网关：</label>
          <div class="col-sm-6">
            <input type="text" class="form-control" name="Gateway" placeholder="请输入备注" value="{{ .idc.Gateway }}">
          </div>
        </div>

        <div class="form-group">
          <label for="Dns" class="col-sm-2 control-label">DNS：</label>
          <div class="col-sm-6">
            <input type="text" class="form-control" name="Dns" placeholder="请输入备注" value="{{ .idc.Dns }}">
          </div>
        </div>

        <div class="form-group">
          <label  class="col-sm-2 control-label" for="IpType">类型：</label>
          <select name="IpType" class="selectpicker" title="请选择类型" data-style="btn-default">
            <option selected value=0>远程卡IP</option>
            <option value=1>业务网卡IP</option>
            <option value=2>数据网卡IP</option>
          </select>
        </div>

        <div class="form-group">
          <label  class="col-sm-2 control-label" for="Status">状态：</label>
          <select name="Status" class="selectpicker" title="请选择状态" data-style="btn-default">
            <option selected value=0>未分配</option>
            <option value=1>已分配</option>
          </select>
        </div>

        <div class="form-group">
          <label for="Comment" class="col-sm-2 control-label">备注：</label>
          <div class="col-sm-6">
            <input type="text" class="form-control" name="Comment" placeholder="请输入备注" value="{{ .idc.Comment }}">
          </div>
        </div>
      </div>

      <div class="modal-footer">
        <button type="button" class="btn btn-default" data-dismiss="modal">关闭</button>
        <button type="submit" class="btn btn-primary">保存</button>
      </div>
    </form>

    <script type="text/javascript">
      $('.selectpicker').selectpicker();
    </script>