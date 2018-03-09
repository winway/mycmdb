    <div class="modal-header">
      <button type="button" class="close" data-dismiss="modal" aria-hidden="true">
        &times;
      </button>
      <h4 class="modal-title" id="myModalLabel">
        IP详情
      </h4>
    </div>
    <form class="form-horizontal" method="post" action="">
      <div class="modal-body">
        <input type="hidden" name="Id" value="{{ .ip.Ip }}">

        <div class="form-group">
          <label  class="col-sm-2 control-label" for="Idc">机房：</label>
          <select name="Idc" class="selectpicker" title="请选择机房" data-style="btn-default" disabled>
            <option selected value={{ .ip.Idc.Name }}>{{ .ip.Idc.Name }}</option>
          </select>
        </div>

        <div class="form-group">
          <label for="Network" class="col-sm-2 control-label">IP/前缀：</label>
          <div class="col-sm-5">
            <input type="text" class="form-control" name="Network" placeholder="请输入IP/前缀" value="{{ .ip.Ip }}/{{ .prefix  }}" disabled>
          </div>
        </div>

        <div class="form-group">
          <label for="SubMask" class="col-sm-2 control-label">子网掩码：</label>
          <div class="col-sm-6">
            <input type="text" class="form-control" name="SubMask" placeholder="请输入备注" value="{{ .ip.SubMask }}" disabled>
          </div>
        </div>

        <div class="form-group">
          <label for="Gateway" class="col-sm-2 control-label">网关：</label>
          <div class="col-sm-6">
            <input type="text" class="form-control" name="Gateway" placeholder="请输入备注" value="{{ .ip.Gateway }}" disabled>
          </div>
        </div>

        <div class="form-group">
          <label for="Dns" class="col-sm-2 control-label">DNS：</label>
          <div class="col-sm-6">
            <input type="text" class="form-control" name="Dns" placeholder="请输入备注" value="{{ .ip.Dns }}" disabled>
          </div>
        </div>

        <div class="form-group">
          <label  class="col-sm-2 control-label" for="IpType">类型：</label>
          <select name="IpType" class="selectpicker" title="请选择类型" data-style="btn-default" disabled>
            {{ if eq 0 .ip.IpType }}
            <option selected value=0>远程卡IP</option>
            <option value=1>业务网卡IP</option>
            <option value=2>数据网卡IP</option>
            {{ else if eq 1 .ip.IpType }}
            <option value=0>远程卡IP</option>
            <option selected value=1>网卡IP</option>
            <option value=2>数据网卡IP</option>
            {{ else if eq 2 .ip.IpType }}
            <option value=0>远程卡IP</option>
            <option value=1>网卡IP</option>
            <option selected value=2>数据网卡IP</option>
            {{ end }}
          </select>
        </div>

        <div class="form-group">
          <label  class="col-sm-2 control-label" for="Status">状态：</label>
          <select name="Status" class="selectpicker" title="请选择状态" data-style="btn-default" disabled>
            {{ if eq 0 .ip.Status }}
            <option selected value=0>未分配</option>
            <option value=1>已分配</option>
            {{ else }}
            <option value=0>未分配</option>
            <option selected value=1>已分配</option>
            {{ end }}
          </select>
        </div>

        <div class="form-group">
          <label for="Comment" class="col-sm-2 control-label">备注：</label>
          <div class="col-sm-6">
            <input type="text" class="form-control" name="Comment" placeholder="请输入备注" value="{{ .ip.Comment }}" disabled>
          </div>
        </div>
      </div>

      <div class="modal-footer">
        <button type="button" class="btn btn-default" data-dismiss="modal">关闭</button>
        <button type="submit" class="btn btn-primary" disabled>保存</button>
      </div>
    </form>

    <script type="text/javascript">
      $('.selectpicker').selectpicker();
    </script>