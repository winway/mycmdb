{{template "base.tpl" .}}

{{define "head"}}
{{end}}

{{define "body"}}

<div class="page-header">
  <h2>IP管理&nbsp;&nbsp;&nbsp;&nbsp;<small>网段{{ .networksCnt }}个 IP地址{{ .ipCnt }}个 已分配IP地址{{ .usedIpCnt }}个</small></h2>
</div>

{{ if .flash.error }}
<div class="alert alert-danger" role="alert">{{ .flash.error }}</div>
{{ else if .flash.notice }}
<div class="alert alert-success" role="alert">{{ .flash.notice }}</div>
{{ end }}

<div>
  <form action="" class="form-inline">
    <div class="form-group">
      <label  class="control-label" for="InputIp">IP:</label>
      <input id="InputIp" class="form-control" placeholder="请输入IP">

      <label  class="control-label" for="selIdc">机房:</label>
      <select id="selIdc" multiple class="selectpicker" title="请选择机房" data-style="btn-default">
        {{range $k, $v := .nameSet}}
        <option value={{ $k }}>{{ $k }}</option>
        {{end}}
      </select>

      <label  class="control-label" for="selNetwork">网段:</label>
      <select id="selNetwork" multiple class="selectpicker" title="请选择网段" data-style="btn-default">
        {{range $index, $elem := .networks}}
        <option value={{ $elem.Network }}>{{ $elem.Network }}</option>
        {{end}}
      </select>

      <label  class="control-label" for="selIpType">类型:</label>
      <select id="selIpType" multiple class="selectpicker" title="请选择类型" data-style="btn-default">
        <option value=0>远程卡IP</option>
        <option value=1>网卡IP</option>
      </select>

      <label  class="control-label" for="selStatus">状态:</label>
      <select id="selStatus" multiple class="selectpicker" title="请选择状态" data-style="btn-default">
        <option value=0>未分配</option>
        <option value=1>已分配</option>
      </select>

      <a onclick="Search()" class="btn btn-primary"><i class="glyphicon glyphicon-search"></i>查询</a>

      <a class="btn btn-primary" href="{{ urlfor "IpController.IpAddPage" }}" data-toggle="modal" data-target="#myModal">
        <i class="glyphicon glyphicon-plus"></i>添加
      </a>

      <!--
      <a class="btn btn-primary" href="{{ urlfor "IpController.IpApplyPage" }}" data-toggle="modal" data-target="#myModal">
        <i class="glyphicon glyphicon-check"></i>申请
      </a>    -->

      <a onclick="Delete('{{ urlfor "IpController.DeleteIp" }}')" class="btn btn-primary"><i class="glyphicon glyphicon-trash"></i>删除</a>
    </div>
  </form>
</div>


<!-- 模态框（Modal） -->
<div class="modal fade" id="myModal" tabindex="-1" role="dialog" aria-labelledby="myModalLabel" aria-hidden="true">
  <div class="modal-dialog">
   <div class="modal-content">
   </div>
 </div><!-- /.modal -->
</div>


<div>
  <table id="myTable"  class="display" cellpadding="0" cellspacing="0" width="100%">
    <thead>
      <tr>
        <th><input type="checkbox" name="checkall" id="checkall"></th>
        <th>IP</th>
        <th>机房</th>
        <th>网段</th>
        <th>类型</th>
        <th>状态</th>
        <th>创建时间</th>
        <th>操作</th>
      </tr>
    </thead>
    <tbody>
    </tbody>
  </table>
</div>


<script type="text/javascript">
  var table;
  $(document).ready(function() {
    table = $('#myTable').dataTable({
    "sPaginationType": "full_numbers", // 分页风格，full_number会把所有页码显示出来（大概是，自己尝试）
    "iDisplayLength": 10,
    "bAutoWidth": false,
    "bLengthChange": false,
    "bFilter": false,
    "bSort": true,
    "oLanguage": {    //下面是一些汉语翻译
      "sSearch": "搜索",
      "sLengthMenu": "每页显示 _MENU_ 条记录",
      "sZeroRecords": "没有检索到数据",
      "sInfo": "显示 _START_ 至 _END_ 条 &nbsp;&nbsp;共 _TOTAL_ 条",
      "sInfoFiltered": "(筛选自 _MAX_ 条数据)",
      "sInfoEmtpy": "没有数据",
      //"sProcessing": "正在加载数据...",
      "sProcessing": "<img src='/static/img/table_loading.gif' />",
      "oPaginate":
      {
        "sFirst": "首页",
        "sPrevious": "前一页",
        "sNext": "后一页",
        "sLast": "末页"
      }
    },
    "bProcessing": true,
    "bServerSide": true,
    "sAjaxSource": "{{ urlfor "IpController.ListIp" }}",
    "aoColumns": [
    {
      "bSortable": false,
      "sClass": "alignRight",
      "mRender": function (data, type, full) {
        return '<input id=' + full['Ip'] + ' value=' + full['Ip'] + ' name="checksingle" type="checkbox">';
      }
    },
    {"mData": 'Ip', "sClass": "alignRight"},
    {"mData": 'Idc.Name', "sClass": "alignRight"},
    {"mData": 'Network', "sClass": "alignRight"},
    {"mData": 'IpType', "sClass": "alignRight"},
    {"mData": 'Status', "sClass": "alignRight"},
    {"mData": 'CreateTime', "sClass": "alignRight"},
    {
      "bSortable": false,
      "sClass": "alignRight",
      "mRender": function (data, type, full) {
        return '<a href="{{ urlfor "IpController.IpDetailPage" ":id" "" }}' + full['Ip'] + '" data-toggle="modal" data-target="#myModal"><i class="glyphicon glyphicon-info-sign" title="详情"></i></a> | <a><i class="glyphicon glyphicon-trash" title="删除" onclick="DeleteOne(\'{{ urlfor "IpController.DeleteIp" }}\',\'' + full['Ip'] + '\')"></i></a>';
      }
    }
    ],
    "aoColumnDefs": [],
    "aaSorting": [[1, "asc"]],
    "fnServerParams": function( aoData )
    {
      var ip = $("#InputIp").val();
      var idc = $("#selIdc").val();
      var network = $("#selNetwork").val();
      var iptype = $("#selIpType").val();
      var status = $("#selStatus").val();
      aoData.push(
        {"name":"filter_ip","value":ip},
        {"name":"filter_idc","value":idc},
        {"name":"filter_network","value":network},
        {"name":"filter_iptype","value":iptype},
        {"name":"filter_status","value":status}
        );
    },
    "fnDrawCallback": function() {
      reloadfunc();
    },
    "fnHeaderCallback": function( nHead, aData, iStart, iEnd, aiDisplay ) {

    }
  });

    $(".alert").delay(4000).slideUp(200, function() {
      $(this).alert('close');
    });

    $("#myTable input[name='checkall']").click(function() {
      $("input[name='checksingle']").prop("checked", this.checked);
    });

    $("#myModal").on("hidden.bs.modal", function() {
      $(this).removeData("bs.modal");
    });

  });
</script>
{{end}}