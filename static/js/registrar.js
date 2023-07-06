$(document).ready(function() {
  // on ready
  
  
});


async function registrarUsuario() {
 let datos = {};
   datos.nombreUsuario= document.getElementById('txtuserName').value;
   datos.mail = document.getElementById('txtMail').value;
   datos.password = document.getElementById('txtPassword').value;
   let repetirPassword = document.getElementById('txtRepetirPassword').value;
   

 if (repetirPassword != datos.password) {
   alert('La contraseña que escribiste es diferente.');
   window.location.href = 'index.html';
 }


const respuesta = await fetch("http://localhost:8080/usuarios/registrar", 
 {
   method: 'POST',
   headers: {
     
     'Content-Type': 'application/json'
   },
   credentials: "include",
   body: JSON.stringify(datos),
 });
 const response = await respuesta.json();
 console.log(data);
   




alert("La cuenta fue creada con exito!");
}



