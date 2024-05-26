document.addEventListener("DOMContentLoaded", function () {
  console.log("Document is Loaded");
  const path = window.location.pathname;
  sendData({
    path: path,
    date: new Date(),
    bucketID: "Something Random"
  });
});

async function sendData(data) {
  var url = "http://localhost:8080/data";

  
  const result = await fetch(url, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(data),
  });

  console.log(data);
}
