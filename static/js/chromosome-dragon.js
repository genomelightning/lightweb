function beginDragon(){
	var chr = $("#chrPicker").val();

	//var srcString = "/static/chromosomes/" + chr + "/chromosome.dzi";
	var srcString = "/static/chromosomes/sm11/chromosome.dzi";
	//try {
		viewer = OpenSeadragon({
      		id: "contentDiv",
			prefixUrl: "/static/js/openseadragon-images/",
        	tileSources: srcString,
			//defaultZoomLevel: 6.1,
			//minZoomLevel: ,
			//maxZoomLevel: 3,
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

}

function dragonWrap(){
//	viewerInputHook = viewer.addViewerInputHook({hooks: [
	//		{tracker: 'viewer', handler: 'clickHandler', hookHandler: onViewerClick}]});
	//function onViewerClick(event) {
	//	event.preventDefaultAction = true;
	//}
	//viewer.destroy();
	viewer.addLayer(OpenSeadragon({
      		id: "contentDiv",
			prefixUrl: "/static/js/openseadragon-images/",
        	tileSources: "/static/chromosomes/sm1/chromosome.dzi",
			//defaultZoomLevel: 6.1,
			//minZoomLevel: ,
			//maxZoomLevel: 3,
			showNavigator: true,
			navigatorPosition: 'BOTTOM_LEFT',
			navigatorHeight: 400,
			navigatorWidth: 80,
			//debugMode: true,
			toolbar: "toolbarDiv",
			//showZoomControl: false
    	}));	
	beginDragon();
}
$(document).ready(beginDragon);

//gigapix and IIPImage
