{{template "base.tpl" .}}

{{define "head"}}
{{end}}

{{define "body"}}

<div class="page-header">
  <h2>代码发布&nbsp;&nbsp;&nbsp;&nbsp;<small>累计发布{{ .doneCnt }}起 处理中{{ .doingCnt }}起</small></h2>
</div>

{{ if .flash.error }}
<div class="alert alert-danger" role="alert">{{ .flash.error }}</div>
{{ else if .flash.notice }}
<div class="alert alert-success" role="alert">{{ .flash.notice }}</div>
{{ end }}

<div>
  <form action="" class="form-inline">
    <div class="form-group">
      <label  class="control-label" for="InputKey">关键字:</label>
      <input id="InputKey" class="form-control" placeholder="请输入关键字">

      <label  class="control-label" for="selStatus">状态:</label>
      <select id="selStatus" multiple class="selectpicker" title="请选择状态" data-style="btn-default">
        <option value=0>进行中</option>
        <option value=1>已完成</option>
      </select>

      <a onclick="Search()" class="btn btn-primary"><i class="glyphicon glyphicon-search"></i>查询</a>

      <a class="btn btn-primary" href="{{ urlfor "ReleaseController.ApplyPage" }}" data-toggle="modal" data-target="#myModal">
        <i class="glyphicon glyphicon-share-alt"></i>发布申请
      </a>
    </div>
  </form>
</div>


<!-- 模态框（Modal） -->
<div class="modal fade" id="myModal" tabindex="-1" role="dialog" aria-labelledby="myModalLabel" aria-hidden="true">
  <div class="modal-dialog  modal-lg">
   <div class="modal-content">
   </div>
 </div><!-- /.modal -->
</div>


<div>
  <table id="myTable"  class="display" cellpadding="0" cellspacing="0" width="100%">
    <thead>
      <tr>
        <th>ID</th>
        <th>主题</th>
        <th>创建者</th>
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
    "sAjaxSource": "{{ urlfor "ReleaseController.List" }}",
    "aoColumns": [
    {"mData": 'Id', "sClass": "alignRight"},
    {"mData": 'Subject', "sClass": "alignRight"},
    {"mData": 'Creator', "sClass": "alignRight"},
    {
      "sClass": "alignRight",
      "mRender": function (data, type, full) {
        if (full["Status"] == 0) {
          return "进行中";
        } else if (full["Status"] == 1) {
          return "已完成";
        }
      }
    },
    {"mData": 'CreateTime', "sClass": "alignRight"},
    {
      "bSortable": false,
      "sClass": "alignRight",
      "mRender": function (data, type, full) {
        return '<a href="{{ urlfor "ReleaseController.ApplyDetailPage" ":id" "" }}' + full['Id'] + '" data-toggle="modal" data-target="#myModal"><i class="glyphicon glyphicon-info-sign" title="详情">流程处理</i></a> | <a><i class="glyphicon glyphicon-trash" title="取消" onclick="DeleteOne(\'{{ urlfor "ReleaseController.Cancel" }}\',\'' + full['Id'] + '\')">取消</i></a>';
      }
    }
    ],
    "aoColumnDefs": [],
    "aaSorting": [[4, "desc"]],
    "fnServerParams": function( aoData )
    {
      var key = $("#InputKey").val();
      var status = $("#selStatus").val();
      aoData.push(
        {"name":"filter_key","value":key},
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