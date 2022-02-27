document.getElementById("dates").classList.add("active")

    var regex = /(?<=e=).*/gm
    var all_inputs = window.location.search;
    var input = all_inputs.match(regex)
    var content = "0; url=/bestdate?trip-start=1970-01-01&trip-end=2022-01-01&artist-date="+input[0]+"&redirected=true"
    if (document.getElementById("meta") != null) {
        document.getElementById("meta").setAttribute("content", content)
    }

    if (input[0].slice(-4,) == "true") {
        document.getElementById("redirected").style.display = "flex"
    }