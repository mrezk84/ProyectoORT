$(document).ready(function() {
  
 });


 async function iniciarSesion() {
  
    let email = document.getElementById('txtEmail').value;
    let password = document.getElementById('txtPassword').value;
     
    const request = await fetch("http://localhost:8080/usuarios/login", {
         method: 'POST',
         headers: {
           'Content-Type': 'application/json',
         },
         credentials: "include",
         body: JSON.stringify({
          email: 'marcos@correo.com',
          password:'123456789'
         }),
       });
       const data = await request.json();
       console.log(data);

}