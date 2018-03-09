  function Search(){
    table.fnDraw();
  }

  function reloadfunc() {
    $("input[name='checksingle']").click(function() {
      var $subs = $("input[name='checksingle']");
      $("#checkall").prop("checked" , $subs.length == $subs.filter(":checked").length ? true : false);
    });
  }

  function DeleteOne(url, id){
    if(confirm("您确认要删除记录" + id + "吗？", function() { }, null)){
      $.ajax({
        type:"POST",
        url:url,
        data:{'ids':[id]},
        success: showmsg
      });
    }
    table.fnDraw();
  }

  function Delete(url){
    var subs = $("input[name='checksingle']");
    if (subs.filter(":checked").length) {
      var ids = getchecked();
      if(confirm("您确认要删除记录" + ids + "吗？", function() { }, null)){
        $.ajax({
          type:"POST",
          url:url,
          data:{'ids':ids},
          success: showmsg
        });
      }
    } else {
      alert("选择项为空！");
    }
    table.fnDraw();
  }

  function getchecked() {
    var l = Array();
    $("input[name='checksingle']").each( function() {
      if($(this).is(":checked")){
        l.push($(this).val());
      }
    });
    return l;
  }

  function showmsg(msg) {
    var obj = eval(msg); //转化为js对象
    alert(obj.info);
  }