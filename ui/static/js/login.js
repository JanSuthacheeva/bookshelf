function togglePWvisible() {
  var x = document.getElementById("password");
  var img = document.getElementById("eyeImg");
  if (x.type === "password") {
    x.type = "text";
    img.src = "../../../static/img/pw-see.svg"
  } else {
    x.type = "password";
    img.src = "../../../static/img/pw-hide.svg"
  }
}

const button = document.getElementById("toggleEye");
button.addEventListener("click", togglePWvisible);
