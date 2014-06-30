function beginDragon(){
	//var chr = $("#chrPicker").val();

	//var srcString = "/static/chromosomes/hu34D5B9.dzi";

	
	var srcString = "/static/chromosomes/all.dzi";
	//try {
		viewer = OpenSeadragon({
      		id: "contentDiv",
			prefixUrl: "/static/js/openseadragon-images/",
        	tileSources: srcString,
			//defaultZoomLevel: 6.1,
			//minZoomLevel: ,
			//maxZoomLevel: 5000,
			visibilityRatio: 0.9,
			showNavigator: true,
			navigatorPosition: 'BOTTOM_LEFT',
			navigatorHeight: 400,
			navigatorWidth: 80,
			//debugMode: true,
			toolbar: "toolbarDiv",
			//showZoomControl: false
    	});
		//Begins an OpenSeadragon.Viewer
		//Adding url hyperlinks would be useful
		//Annotation of genes, etc	

	//}
	//catch (){}

	viewerInputHook = viewer.addViewerInputHook({hooks: [
			{tracker: 'viewer', handler: 'clickHandler', hookHandler: onViewerClick}]});
	function onViewerClick(event) {
		event.preventDefaultAction = true;
		console.log(viewer.viewport.getZoom());
		console.log(viewer.viewport.getMaxZoom());
	}

}

$(document).ready(beginDragon);

//gigapix and IIPImage
