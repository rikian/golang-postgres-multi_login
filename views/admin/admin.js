function qs(elm) {
  return document.querySelector(elm);
}

const xml = new XMLHttpRequest();

const btnLogout = qs(".btnLogout");
btnLogout.addEventListener("click", function (e) {
  xml.withCredentials = true;
  xml.open("POST", `${window.origin}/logout`, true);
  xml.setRequestHeader("content-type", "application/json");
  xml.send(JSON.stringify({ method: "logout" }));
  // console.log(e);
  return;
});

xml.onload = function (e) {
  console.log(JSON.parse(xml.response));
  window.location.reload();
  return;
};
