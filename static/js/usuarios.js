// Call the dataTables jQuery plugin
$(document).ready(function() {
  cargarUsuarios();

$('#usuarios').DataTable();


});

function actualizarCedulaDelUsuario() {
  document.getElementById('txt-nombre-usuario').outerHTML = localStorage.nombre
}
async function cargarRoles() {
  const request = await fetch('http://localhost:8080/usuarios', {
    method: 'GET',
    headers: getHeaders()
  });
  const roles = await request.json();
  if(roles.request==200){
   botonRol = '<a href="roles.html" onclick="asingarRol()" class="btn btn-success btn-icon-split"><span class="icon text-white-50"><i class="fas fa-check"></i></span><span class="text">Asignar</span></a>';  
  }
}

async function cargarUsuarios() {
  const request = await fetch("http://localhost:8080/usuarios", {
    method: 'GET',
  })
  let response = await request.json()
  console.log(response)
  if (request.ok) {
            let text = ``
            response.forEach(usuario => {
                text +=
                    `
                     <tr>
                        <th>${usuario.id}</th>
                        <th>${usuario.email}</th>
                        <th>${usuario.name}</th>
                    </tr>
                    `
            });
            document.getElementById("usuariosTBody").innerHTML = text;
  }
document.querySelector('#usuarios tbody').outerHTML = listadoHtml;

}


function getHeaders() {
  return {
   'Accept': 'application/json',
   'Content-Type': 'application/json'
     
 };
}

async function eliminarUsuario(id) {

if (!confirm('¿Desea eliminar este usuario?')) {
  return;
}

const request = await fetch('http://localhost:8080/usuarios' + id, {
  method: 'DELETE',
  headers: getHeaders()
});
location.reload();
}
async function editarUsuario(id) {

  if (!confirm('¿Desea editar el usuario?')) {
    return;
  }
  
  const request = await fetch('api/usuarios/' + id, {
    method: 'POST',
    headers: getHeaders()
  });
  
}

async function asignaroRol(nombre) {

  if (!confirm('¿Desea asingar este rol al usuario?')) {
     return;
   }
 const request = await fetch('http://localhost:8080/usuarios' + nombre, {
   method: 'POST',
   headers: getHeaders()
 });

 
 
 
} 

