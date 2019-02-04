$('.nav_websites').click(function() {
	$('#content').load('websites.php');
	$('#content').attr('class','isotope-container');
	return false;
});