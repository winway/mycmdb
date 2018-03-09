      <nav class="navbar navbar-default">
        <div class="container-fluid">
          <div class="navbar-header">
            <button type="button" id="sidebarCollapse" class="btn btn-info navbar-btn">
              <i class="glyphicon glyphicon-align-left"></i>
              <span></span>
            </button>
          </div>

          <div class="collapse navbar-collapse" id="bs-example-navbar-collapse-1">
            <ul class="nav navbar-nav navbar-right">
              <span class="badge"><big>你好, {{ .username }}</big></span>
              <!-- Split button -->
              <div class="btn-group navbar-btn">
                <a href="{{ urlfor "UserController.Logout" }}" type="button" class="btn btn-default">注销</a>
                <button type="button" class="btn btn-default dropdown-toggle" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
                  <span class="caret"></span>
                  <span class="sr-only">Toggle Dropdown</span>
                </button>
                <ul class="dropdown-menu">
                  <li><a href="#">设置</a></li>
                  <li role="separator" class="divider"></li>
                  <li><a href="#">修改密码</a></li>
                </ul>
              </div>
            </ul>
          </div>
        </div>
      </nav>