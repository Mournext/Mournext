{{define "headerb"}}
<!DOCTYPE html>
<html lang="en">
<head>
  
  <!-- start: Meta -->
  <meta charset="utf-8">
  <meta name="description" content="Tasks">
  <meta name="author" content="MTY">
  <meta name="keyword" content="Task, Metro UI, Bootstrap, Admin, Template, Theme, Responsive, Fluid, Retina">
  <!-- end: Meta -->
  
  <!-- start: Mobile Specific -->
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <!-- end: Mobile Specific -->
  
  <!-- start: CSS -->
  <link id="bootstrap-style" href="/static/css/bootstrap.min.css" rel="stylesheet">
  <link href="/static/css/bootstrap-responsive.min.css" rel="stylesheet">
  <link id="base-style" href="/static/css/style.css" rel="stylesheet">
  <link id="base-style-responsive" href="/static/css/style-responsive.css" rel="stylesheet">
  <link href='http://fonts.googleapis.com/css?family=Open+Sans:300italic,400italic,600italic,700italic,800italic,400,300,600,700,800&subset=latin,cyrillic-ext,latin-ext' rel='stylesheet' type='text/css'>
  <!-- end: CSS -->
  

  <!-- The HTML5 shim, for IE6-8 support of HTML5 elements -->
  <!--[if lt IE 9]>
      <script src="http://html5shim.googlecode.com/svn/trunk/html5.js"></script>
    <link id="ie-style" href="css/ie.css" rel="stylesheet">
  <![endif]-->
  
  <!--[if IE 9]>
    <link id="ie9style" href="css/ie9.css" rel="stylesheet">
  <![endif]-->
    
  <!-- start: Favicon -->
  <link rel="shortcut icon" href="/static/img/favicon.png">
  <!-- end: Favicon -->
     <!-- start: JavaScript-->       
    <script src="/static/js/jquery-1.9.1.min.js"></script>
  <script src="/static/js/jquery-migrate-1.0.0.min.js"></script>
  
    <script src="/static/js/jquery-ui-1.10.0.custom.min.js"></script>
  
    <script src="/static/js/jquery.ui.touch-punch.js"></script>
  
    <script src="/static/js/modernizr.js"></script>
  
    <script src="/static/js/bootstrap.min.js"></script>
  
    <script src="/static/js/jquery.cookie.js"></script>
  
    <script src='/static/js/fullcalendar.min.js'></script>
  
    <script src='/static/js/jquery.dataTables.min.js'></script>

    <script src="/static/js/excanvas.js"></script>
  <script src="/static/js/jquery.flot.js"></script>
  <script src="/static/js/jquery.flot.pie.js"></script>
  <script src="/static/js/jquery.flot.stack.js"></script>
  <script src="/static/js/jquery.flot.resize.min.js"></script>
  
    <script src="/static/js/jquery.chosen.min.js"></script>
  
    <script src="/static/js/jquery.uniform.min.js"></script>
    
    <script src="/static/js/jquery.cleditor.min.js"></script>
  
    <script src="/static/js/jquery.noty.js"></script>
  
    <script src="/static/js/jquery.elfinder.min.js"></script>
  
    <script src="/static/js/jquery.raty.min.js"></script>
  
    <script src="/static/js/jquery.iphone.toggle.js"></script>
  
    <script src="/static/js/jquery.uploadify-3.1.min.js"></script>
  
    <script src="/static/js/jquery.gritter.min.js"></script>
  
    <script src="/static/js/jquery.imagesloaded.js"></script>
  
    <script src="/static/js/jquery.masonry.min.js"></script>
  
    <script src="/static/js/jquery.knob.modified.js"></script>
  
    <script src="/static/js/jquery.sparkline.min.js"></script>
  
    <script src="/static/js/counter.js"></script>
  
    <script src="/static/js/retina.js"></script>

    <script src="/static/js/custom.js"></script>
    <!-- end: JavaScript-->
    
</head>

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