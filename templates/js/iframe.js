"use strict";
(function () {
    $('#myIframe').load(function(){
        var iframe = document.getElementById("myIframe");
        var footer = iframe.getElementsById("footer")
        footer.style.display = "none";
        console.log("I`m here")
    });
    function hideIframeFooter() {
        var iframe = document.getElementById("myIframe");
        var footer = iframe.getElementsById("footer")
        footer.style.display = "none";
        console.log("I`m here 2")
    }

    setInterval(hideIframeFooter,1000)
    console.log("fffffffffffffffffffffff")

}());

