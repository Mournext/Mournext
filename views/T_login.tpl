{{define "login"}}
<body>
    <!-- start: Header -->
  <div class="navbar">
    <div class="navbar-inner">
      <div class="container-fluid">
        <a class="btn btn-navbar" data-toggle="collapse" data-target=".top-nav.nav-collapse,.sidebar-nav.nav-collapse">
          <span class="icon-bar"></span>
          <span class="icon-bar"></span>
          <span class="icon-bar"></span>
        </a>
        <a class="brand" href="/"><span>Task</span></a>
                
        <!-- start: Header Menu -->
        <div class="nav-no-collapse header-nav">
          <ul class="nav pull-right"> 
          {{if .IsLogin}}             
            <li><a href="/login?exit=true"><i class="halflings-icon white off"></i> Logout</a></li>
            {{else}}
            <li><a href="/login"><i class="halflings-icon white user"></i> Login</a></li>
            {{end}}             
          </ul>
        </div>
        <!-- end: Header Menu -->
        
      </div>
    </div>
  </div>
  <!-- start: Header -->
  
    <div class="container-fluid-full">
    <div class="row-fluid">
        
      <!-- start: Main Menu -->
      <div id="sidebar-left" class="span2">
        <div class="nav-collapse sidebar-nav">
          <ul class="nav nav-tabs nav-stacked main-menu">
            
            <li><a href="/task"><i class="icon-list"></i><span class="hidden-tablet"> Task</span></a></li>         
            <li>
              <a href="/category"><i class="icon-folder-close-alt"></i><span class="hidden-tablet"> Category</span></a>
            </li>
            <li><a href="/icon"><i class="icon-star"></i><span class="hidden-tablet"> About Icons</span></a></li>
            <li><a href="/login"><i class="icon-lock"></i><span class="hidden-tablet"> Login Page</span></a></li>
          </ul>
        </div>
      </div>
      <!-- end: Main Menu -->
      
      <noscript>
        <div class="alert alert-block span10">
          <h4 class="alert-heading">Warning!</h4>
          <p>You need to have <a href="http://en.wikipedia.org/wiki/JavaScript" target="_blank">JavaScript</a> enabled to use this site.</p>
        </div>
      </noscript>    
  
{{end}}


 


