1. Add amplify for nginx monitoring
https://amplify.nginx.com/signup/

8. no magic numbers in server. use constants

9. Need to change to 3.6 mongodb version

11. rewrite the whole docker shit

12. date object should in vue should be a single object of multiple properties

<script>
    function hello() {
      console.log("Hello");

      var http = new XMLHttpRequest();
      var url = "http://localhost:8080/";
      var params = "foo=bar";

      http.open("POST", url);
      http.setRequestHeader("Content-type", "application/x-www-form-urlencoded");

      http.onreadystatechange = function () { //Call a function when the state changes.
        if (http.readyState == 4 && http.status == 200) {
          console.log(JSON.parse(http.responseText));
        }
      }
      http.send(params);
    }
</script> 