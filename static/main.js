$(function(){
    $.get("/ping", function(data){
        if(data.error == "true"){
            $("#results").prepend("<div class='alert alert-danger'><strong>Error!</strong> "+ data.message +"</div>");
        }
    }, "json");

    var getalltrips = $.get("/allTrips", function(data) {
        var arr = data.result;
        arr.splice(0,1);
        var count = arr.length / 5;
        var posts = new Array();
        for (var i = 0; i < count; i++) {
            posts.push(arr.splice(0, 5))
        }
        for (var i = 0; i < posts.length; i++) {
            var post = document.createElement("div");
            post.classList.add("post");
            post.classList.add("media");

            var medialeft = document.createElement("div");
            medialeft.classList.add("media-left");

            var mediaobject = document.createElement("img");
            mediaobject.classList.add("media-object");
            mediaobject.src = posts[i][1];

            var mediabody = document.createElement("div");
            mediabody.classList.add("media-body");

            var mediaheading = document.createElement("h4");
            mediabody.classList.add("media-heading");
            var name = document.createElement("span");
            name.classList.add("username");
            name.innerHTML = posts[i][0];
            var title = document.createElement("span");
            title.classList.add("triptitle");
            title.classList.add("title" + i);
            title.innerHTML = posts[i][2];
            mediaheading.appendChild(name);
            mediaheading.appendChild(document.createTextNode(" shared "));
            mediaheading.appendChild(title);
            
            var span = document.createElement("span");
            span.classList.add("date");
            var date = posts[i][4];
            span.innerHTML = moment(date).fromNow();

            var p = document.createElement("p");
            p.innerHTML = posts[i][3];

            medialeft.appendChild(mediaobject);
            mediabody.appendChild(mediaheading);
            mediaheading.appendChild(span);
            mediabody.appendChild(p);
            post.appendChild(medialeft);
            post.appendChild(mediabody);
            document.querySelector(".feeds").appendChild(post);
        }
    }, "json");

    $(".toggle").click(function(){
        $(".form-container").slideToggle();
    });

    $("#submittrippoint").click(function(){
        submittrippoint();
        $("#date").val("");
        $("#trippointdescription").val("");
        $("#address1").val("");
        $("#city").val("");
        $("#country").val("");
        $("#transportationtype").val("");
        $("#transportationcost").val("");
        $("#transportation").val("");

        $(".toggle").click();
        $(".toggle").click();
    })

    $("#submittrip").click(function(){
        submittrip();
        changeview();
    })

    $("#finish").click(function() {
        $("#trip-name").val("");
        $("#trip-description").val("");
        changeview();
    })

    setTimeout(function() {
        $(".triptitle").click(function() {
            gettripinfo($(this).text());
        })
    }, 2000);

    function gettripinfo(title) {
        $.post("/trip", {triptitle: title})
        .done(function(data) {
            document.querySelector(".view").innerHTML="";
            var arr = data.result;
            arr.splice(0,1);
            var count = arr.length / 9;
            var points = new Array();
            for (var i = 0; i < count; i++) {
                points.push(arr.splice(0, 9))
            }
            console.log(points);

            var ul = document.createElement("ul");
            ul.classList.add("tripview");
            ul.classList.add("list-group");

            var h4 = document.createElement("h4");
            h4.innerHTML = points[0][2];
            var h6 = document.createElement("h6");
            h6.innerHTML = points[0][3];

            var liitem = document.createElement("li");
            liitem.classList.add("list-group-item");
            liitem.appendChild(h4);
            liitem.appendChild(h6);
            ul.appendChild(liitem);

            for (var i = 0; i < points.length; i++) {
                var li = document.createElement("li");
                li.classList.add("list-group-item");
                
                var num = document.createElement("span");
                num.classList.add("num");
                num.innerHTML = i + 1;

                var date = document.createElement("div");
                date.classList.add("date");
                date.innerHTML = moment(points[i][4]).format("Do MMMM YYYY | h:mm a");

                var country = document.createElement("span");
                country.innerHTML = points[i][6];
                country.classList.add("country")

                var city = document.createElement("span");
                city.innerHTML = points[i][7];
                city.classList.add("city");

                var transportation = document.createElement("span");
                transportation.innerHTML = points[i][8];
                transportation.classList.add("transportation");

                var p = document.createElement("p");
                p.innerHTML = points[i][5];

                var div = document.createElement("div");
                div.appendChild(document.createTextNode("Went to "));
                div.appendChild(city);
                div.appendChild(document.createTextNode(", "));
                div.appendChild(country);
                div.appendChild(document.createTextNode(" via "));
                div.appendChild(transportation);
                div.appendChild(p);

                li.appendChild(date);
                li.appendChild(div);
                ul.appendChild(li);
            }
            document.querySelector(".view").appendChild(ul);
        })

    }

    function changeview() {
        $(".toggle").click();
        $(".trip-point").toggle();
        $(".trip").toggle();
        $("#submittrippoint").toggle();
        $("#finish").toggle();
        $("#submittrip").toggle();
        $(".toggle").click();
    }

    function submittrip() {
        $.post("/addtrip", {name: $("#trip-name").val(), description: $("#trip-description").val()})
        .done(function(data){
            console.log(data);
        });
    }

    function submittrippoint() {
        $.post("/addtrippoint", 
            {
            date: $("#date").val(), 
            trippointdescription: $("#trippointdescription").val(), 
            address1: $("#address1").val(), 
            city: $("#city").val(), 
            country: $("#country").val(), 
            transportationtype: $("#transportationtype").val(),
            transportationcost: $("#transportationcost").val(),
            transportation: $("#transportation").val()
        })
            .done(function(data){
            console.log(data);
        });
    }

    /*
    $("#login").click(function(){
      $.post("/login", {username: $("#username").val(), password: $("#password").val()})
        .done(function(data){
          if(data){
            console.log(data)
            $("#result").text("Logged in as: "+data.username);
          } else {
            console.log("Failed to log in!")
            $("#result").text("Username / pasword combination invalid!");
          }
        });
    });
    */
})
