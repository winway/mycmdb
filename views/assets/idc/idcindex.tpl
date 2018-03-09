{{template "base.tpl" .}}

{{define "head"}}
{{end}}

{{define "body"}}

<div class="page-header">
  <h2>机房管理&nbsp;&nbsp;&nbsp;&nbsp;<small>机房{{ .cnt }}个</small></h2>
</div>

{{ if .flash.error }}
<div class="alert alert-danger" role="alert">{{ .flash.error }}</div>
{{ else if .flash.notice }}
<div class="alert alert-success" role="alert">{{ .flash.notice }}</div>
{{ end }}

<div>
  <form action="" class="form-inline">
    <div class="form-group">
      <label  class="control-label" for="selName">名称:</label>
      <select id="selName" multiple class="selectpicker" title="请选择名称" data-style="btn-default">
        {{range $index, $elem := .names}}
        <option value={{ $elem.Name }}>{{ $elem.Name }}</option>
        {{end}}
      </select>

      <label  class="control-label" for="selOperator">运营商:</label>
      <select id="selOperator" multiple class="selectpicker" title="请选择运营商" data-style="btn-default">
        {{range $index, $elem := .operators}}
        <option value={{ $elem.Operator }}>{{ $elem.Operator }}</option>
        {{end}}
      </select>

      <label  class="control-label" for="selLinkman">联系人:</label>
      <select id="selLinkman" multiple class="selectpicker" title="请选择联系人" data-style="btn-default">
        {{range $index, $elem := .linkmans}}
        <option value={{ $elem.Linkman }}>{{ $elem.Linkman }}</option>
        {{end}}
      </select>

      <a onclick="Search()" class="btn btn-primary"><i class="glyphicon glyphicon-search"></i>查询</a>

      <a class="btn btn-primary" href="{{ urlfor "IdcController.IdcAddPage" }}" data-toggle="modal" data-target="#myModal">
        <i class="glyphicon glyphicon-plus"></i>添加
      </a>

      <a onclick="Delete('{{ urlfor "IdcController.DeleteIdc" }}')" class="btn btn-primary"><i class="glyphicon glyphicon-trash"></i>删除</a>
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
        <th>ID</th>
        <th>名称</th>
        <th>地址</th>
        <th>运营商</th>
        <th>联系人</th>
        <th>联系电话</th>
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
    "sAjaxSource": "{{ urlfor "IdcController.ListIdc" }}",
    "aoColumns": [
    {
      "bSortable": false,
      "sClass": "alignRight",
      "mRender": function (data, type, full) {
        return '<input id=' + full['Id'] + ' value=' + full['Id'] + ' name="checksingle" type="checkbox">';
      }
    },
    {"mData": 'Id', "sClass": "alignRight"},
    {"mData": 'Name', "sClass": "alignRight"},
    {"mData": 'Address', "sClass": "alignRight"},
    {"mData": 'Operator', "sClass": "alignRight"},
    {"mData": 'Linkman', "sClass": "alignRight"},
    {"mData": 'Phone', "sClass": "alignRight"},
    {"mData": 'CreateTime', "sClass": "alignRight"},
    {
      "bSortable": false,
      "sClass": "alignRight",
      "mRender": function (data, type, full) {
        return '<a href="{{ urlfor "IdcController.IdcDetailPage" ":id" "" }}' + full['Id'] + '" data-toggle="modal" data-target="#myModal"><i class="glyphicon glyphicon-info-sign" title="详情"></i></a> | <a href="{{ urlfor "IdcController.IdcEditPage" ":id" "" }}' + full['Id'] + '" data-toggle="modal" data-target="#myModal"><i class="glyphicon glyphicon-edit" title="编辑"></i></a> | <a><i class="glyphicon glyphicon-trash" title="删除" onclick="DeleteOne(\'{{ urlfor "IdcController.DeleteIdc" }}\',' + full['Id'] + ')"></i></a>';
      }
    }
    ],
    "aoColumnDefs": [],
    "aaSorting": [[1, "asc"]],
    "fnServerParams": function( aoData )
    {
      var name = $("#selName").val();
      var operator = $("#selOperator").val();
      var linkman = $("#selLinkman").val();
      aoData.push(
        {"name":"filter_name","value":name},
        {"name":"filter_operator","value":operator},
        {"name":"filter_linkman","value":linkman}
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