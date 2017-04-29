(function() {
  var directionsDisplay;
  var directionsService;
  var markersArray = [];
  var waypoints = [];
  var infowindow;
  var startloc;
  var service;
  var item = document.getElementById("options");
  var k = 0;

  // Route will create new direction service, and also it will remove
  // all markers that exist if a previous route has been constructed.
  function Route() {
    var stepDisplay = new google.maps.InfoWindow;
    if(directionsDisplay != null) {
      directionsDisplay.setMap(null);
    }
    directionsService = new google.maps.DirectionsService;
    directionsDisplay = new google.maps.DirectionsRenderer;
    directionsDisplay.setMap(map);

    for(var i = 0; i < markersArray.length; i++) {
      markersArray[i].setMap(null);
    }
    markersArray = [];

    calculateAndDisplayRoute(directionsService, directionsDisplay, markersArray, stepDisplay, map);
  }
  window.Route = Route;

  // calculateAndDisplayRoute will route between point A and B and will
  // call the functionality to add the markers along the route.
  function calculateAndDisplayRoute(directionsService, directionsDisplay, markerArray, stepDisplay, map) {
    directionsService.route({
      origin: document.getElementById('start').value,
      destination: document.getElementById('end').value,
      // The default is driving, but this provides clarity and possible later flexibility
      travelMode: 'DRIVING'
    }, function(response, status) {
      if (status === 'OK') {
        directionsDisplay.setDirections(response);
        showSteps(response, markerArray, stepDisplay, map);
      } else {
        window.alert('Directions request failed due to ' + status);
      }
    });
  }

  //showSets will go through each directionResult instreaction and do the
  //nearby search at those points.  This will be changed to be equally
  //incremental points along the route instead of direction changes.
  // google.maps.geometry.spherical.computeDistanceBetween (latLngA, latLngB);

  function showSteps(directionResult, markerArray, stepDisplay, map) {

    var myRoute = directionResult.routes[0].legs[0];
    var start = myRoute.steps[0].start_location;
    var end = myRoute.steps[myRoute.steps.length-1].end_location;

    infowindow = new google.maps.InfoWindow();

    //item = document.getElementById("options");
    service = new google.maps.places.PlacesService(map);
    ;
    for(k; k < myRoute.steps.length-1; k++) {

      startloc = myRoute.steps[k].start_location;
      var endpos = myRoute.steps[k+1].start_location;


      if(shortDistance(startloc, endpos, start, end)) {
        setTimeout(function(k){
          return function() {
            service.nearbySearch({
              location: startloc,
              radius: 15000,
              type: item.options[item.selectedIndex].text
            }, callback);
          };
        }(k), 1000*k);
        continue;

      }

      while(longDistance(startloc, endpos, start, end)) {
        setTimeout(nearbyClosure(startloc, service, item, callback), 1000*k);
        var heading = google.maps.geometry.spherical.computeHeading(startloc, endpos);
        startloc = google.maps.geometry.spherical.computeOffset(startloc, 32000, heading);
      }

    }

  }

  //callback function to createMarkers for the nearby search
  function callback(results, status) {
    if (status == google.maps.places.PlacesServiceStatus.OK) {
      for (var i = 0; i < results.length; i++) {
        createMarker(results[i]);
      }
    }
    if (status == google.maps.GeocoderStatus.OVER_QUERY_LIMIT) {
      console.log("over");
    }
  }

  // createMarker takes in a place and creates the infowindow markersArray.
  // places markers on the map and adds to the markersArray.
  function createMarker(place) {
    var placelat = place.geometry.location.lat();
    var placelong = place.geometry.location.lng();
    if(place.rating != null && place.photos != null) {
      var marker = new google.maps.Marker({
        map: map,
        position: place.geometry.location
        //icon: place.photos[0].getUrl({'maxWidth': 35, 'maxHeight': 35})
      });
      //console.log(marker.position.toString());
      //photUrl is the location of the photo to be placed in the
      var photoUrl = place.photos[0].getUrl({'maxWidth': 250, 'maxHeight': 250});
      var contentString =  "<img src=" + photoUrl + ">"+"<br />"+
                            place.name + "<br />Rating: " +
                            place.rating + "<br/ >" +
                            '<button onclick="addWaypoint('+ placelat +',' + placelong + ')">Add to Route</button>';

      // addListener for click marker
      google.maps.event.addListener(marker, 'click', function() {
        infowindow.setContent(contentString);
        infowindow.open(map, this);
      });

      markersArray.push(marker);
    }
  }

  function addWaypoint(lat, long) {
      waypoints.push({
        location: new google.maps.LatLng(lat, long),
        stopover: true
      });
      console.log(waypoints);
  }
  window.addWaypoint = addWaypoint;

  function genFinalRoute() {
    directionsService = new google.maps.DirectionsService;
    directionsDisplay = new google.maps.DirectionsRenderer;
    if(directionsDisplay != null) {
      directionsDisplay.setMap(null);
    }
    for(var i = 0; i < markersArray.length; i++) {
      markersArray[i].setMap(null);
    }

    directionsDisplay.setMap(map);



    directionsService.route({
      origin: document.getElementById('start').value,
      destination: document.getElementById('end').value,
      waypoints: waypoints,
      optimizeWaypoints: true,
      // The default is driving, but this provides clarity and possible later flexibility
      travelMode: 'DRIVING'
    }, function(response, status) {
      if (status === 'OK') {
        directionsDisplay.setDirections(response);
      } else {
        window.alert('Directions request failed due to ' + status);
      }
    });
  }
  window.genFinalRoute = genFinalRoute;


  function nearbyClosure(startloc, service, item, callback) {
    return function() {
      service.nearbySearch({
        location: startloc,
        radius: 15000,
        type: item.options[item.selectedIndex].text
      }, callback);
    };
  }

  // WTF at these variable names?
  function shortDistance(startloc, endpos, start, end) {
    return (google.maps.geometry.spherical.computeDistanceBetween(startloc, endpos) < 32000 &&
    google.maps.geometry.spherical.computeDistanceBetween(startloc, start) > 300000 &&
    google.maps.geometry.spherical.computeDistanceBetween(startloc, end) > 300000);
  }

  // again, wtf?
  function longDistance(startloc, endpos, start, end) {
    return (google.maps.geometry.spherical.computeDistanceBetween(startloc, endpos) > 32000 &&
    google.maps.geometry.spherical.computeDistanceBetween(startloc, start) > 300000 &&
    google.maps.geometry.spherical.computeDistanceBetween(startloc, end) > 300000);
  }
}());
