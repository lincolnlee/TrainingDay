<!DOCTYPE html>
<html lang="en">
<head>
	
	<!-- start: Meta -->
	<meta charset="utf-8">
	<title>Optimus Dashboard</title>
	<meta name="description" content="Optimus Dashboard Bootstrap Admin Template.">
	<meta name="author" content="Łukasz Holeczek">
	<!-- end: Meta -->
	
	<!-- start: Mobile Specific -->
	<meta name="viewport" content="width=device-width, initial-scale=1">
	<!-- end: Mobile Specific -->

	<!-- start: CSS -->
	<link id="bootstrap-style" href="static/css/bootstrap.css" rel="stylesheet">
	<link href="static/css/bootstrap-responsive.min.css" rel="stylesheet">
	<link id="base-style" href="static/css/style.css" rel="stylesheet">
	<link id="base-style-responsive" href="static/css/style-responsive.css" rel="stylesheet">
	<!-- end: CSS -->

	<!-- The HTML5 shim, for IE6-8 support of HTML5 elements -->
	<!--[if lt IE 9]>
	  <script src="http://html5shim.googlecode.com/svn/trunk/html5.js"></script>
	<![endif]-->

	<!-- start: Favicon -->
	<link rel="shortcut icon" href="static/img/favicon.ico">
	<!-- end: Favicon -->
		
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
				<a class="brand" href="default"> <img alt="Optimus Dashboard" src="static/img/logo20.png" /> <span>RSS Reader</span></a>
				
				
			</div>
		</div>
	</div>
	<div id="under-header"></div>
	<!-- start: Header -->
	
		<div class="container-fluid">
		<div class="row-fluid">
				
			<!-- start: Main Menu -->
			<div class="span2 main-menu-span">
				<div class="nav-collapse sidebar-nav">
					<ul class="nav nav-tabs nav-stacked main-menu">
						<li class="nav-header hidden-tablet">Navigation</li>
						<li><a href="default"><i class="icon-align-justify"></i><span class="hidden-tablet"> Read Info</span></a></li>
						<li><a href="setting"><i class="icon-edit"></i><span class="hidden-tablet"> Setting</span></a></li>
						<li><a href="index"><i class="icon-lock"></i><span class="hidden-tablet"> Login</span></a></li>
					</ul>
				</div><!--/.well -->
			</div><!--/span-->
			<!-- end: Main Menu -->
			
			<noscript>
				<div class="alert alert-block span10">
					<h4 class="alert-heading">Warning!</h4>
					<p>You need to have <a href="http://en.wikipedia.org/wiki/JavaScript" target="_blank">JavaScript</a> enabled to use this site.</p>
				</div>
			</noscript>
			
			<div id="content" class="span10">
			<!-- start: Content -->
			

			<div>
				<ul class="breadcrumb">
					<li>
						<a href="#">Home</a> <span class="divider">/</span>
					</li>
					<li>
						<a href="#">Read Info</a>
					</li>
				</ul>
			</div>
			

			<div class="row-fluid sortable">
				<div class="box span12">
					<div class="box-header">
						<h2><i class="icon-align-justify"></i><span class="break"></span>Info</h2>
						<div class="box-icon">
							<a href="#" class="btn-setting"><i class="icon-wrench"></i></a>
							<a href="#" class="btn-minimize"><i class="icon-chevron-up"></i></a>
							<a href="#" class="btn-close"><i class="icon-remove"></i></a>
						</div>
					</div>
					<div class="box-content">
						<table class="table">
							  <thead>
								  <tr>
									  <th>Title</th>
									  <th>Time</th>
									  <th>Src Name</th>                                        
								  </tr>
							  </thead>   
							  <tbody>
							  	{{range .RssContentLs}}
								<tr>
									<td><a href="{{.Alink}}" title="{{.Adesc}}" target="_blank">{{.Atitle}}</a></td>
									<td class="center">{{.Atime}}</td>
									<td class="center">{{.SrcName}}</td>                                      
								</tr>  
								{{end}}                                 
							  </tbody>
						 </table>  
						 <div class="pagination pagination-centered">
						  <ul>
						  	{{range .Page}}
							<li {{if .IsCurrent}}class="active"{{end}}><a href="default?page={{.PageIdx}}">{{.PageIdx}}</a></li>
							{{end}}
						  </ul>
						</div>     
					</div>
				</div><!--/span-->
				
				
			</div><!--/row-->
			
			
    
					<!-- end: Content -->
		</div><!--/#content.span10-->
		</div><!--/fluid-row-->
				
		<div class="modal hide fade" id="myModal">
			<div class="modal-header">
				<button type="button" class="close" data-dismiss="modal">×</button>
				<h3>Settings</h3>
			</div>
			<div class="modal-body">
				<p>Here settings can be configured...</p>
			</div>
			<div class="modal-footer">
				<a href="#" class="btn" data-dismiss="modal">Close</a>
				<a href="#" class="btn btn-primary">Save changes</a>
			</div>
		</div>
		
		<div class="clearfix"></div>
		<hr>
		
		<footer>
			<p class="pull-left">&copy; <a href="" target="_blank">creativeLabs</a> 2013</p>
			<p class="pull-right">Powered by: <a href="#">Optimus Dashboard</a></p>
		</footer>
				
	</div><!--/.fluid-container-->

	<!-- start: JavaScript-->

		<script src="static/js/jquery-1.9.1.min.js"></script>
		<script src="static/js/jquery-migrate-1.0.0.min.js"></script>
		<script src="static/js/jquery-ui-1.10.0.custom.min.js"></script>
	
		<script src="static/js/bootstrap.js"></script>
	
		<script src="static/js/jquery.cookie.js"></script>
	
		<script src='static/js/fullcalendar.min.js'></script>
	
		<script src='static/js/jquery.dataTables.min.js'></script>

		<script src="static/js/excanvas.js"></script>
		<script src="static/js/jquery.flot.min.js"></script>
		<script src="static/js/jquery.flot.pie.min.js"></script>
		<script src="static/js/jquery.flot.stack.js"></script>
		<script src="static/js/jquery.flot.resize.min.js"></script>
	
		<script src="static/js/jquery.chosen.min.js"></script>
	
		<script src="static/js/jquery.uniform.min.js"></script>
		
		<script src="static/js/jquery.cleditor.min.js"></script>
	
		<script src="static/js/jquery.noty.js"></script>
	
		<script src="static/js/jquery.elfinder.min.js"></script>
	
		<script src="static/js/jquery.raty.min.js"></script>
	
		<script src="static/js/jquery.iphone.toggle.js"></script>
	
		<script src="static/js/jquery.uploadify-3.1.min.js"></script>
	
		<script src="static/js/jquery.gritter.min.js"></script>
	
		<script src="static/js/jquery.imagesloaded.js"></script>
	
		<script src="static/js/jquery.masonry.min.js"></script>
	
		<script src="static/js/custom.js"></script>

		<!-- end: JavaScript-->
	
</body>
</html>
