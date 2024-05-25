



document.addEventListener("DOMContentLoaded" ,function (){


    console.log("Document is Loaded");

    sendData()
    
    
})

 async function sendData(data) {
    
    var url="http://localhost:8080/data"
    
    var params={foo:"bar"}
    
   const result= await fetch(url,{
    
        method:'POST',
        headers:{
            'Content-Type':'application/json',
    
        },
        body:JSON.stringify(data)
    
    })
}

