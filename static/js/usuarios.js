// Call the dataTables jQuery plugin
$(document).ready(function() {
  cargarUsuarios();
  actualizarEmailDelUsuario()
  Paginar("usuarios");


});

function actualizarEmailDelUsuario() {
  document.getElementById('txt-nombre-usuario').outerHTML = localStorage.name
}
async function cargarRoles() {
  const request = await fetch('http://localhost:5000/usuarios', {
    method: 'GET',
    headers: getHeaders()
  });
  const roles = await request.json();
  if(roles.request==200){
   botonRol = '<a href="roles.html" onclick="asingarRol()" class="btn btn-success btn-icon-split"><span class="icon text-white-50"><i class="fas fa-check"></i></span><span class="text">Asignar</span></a>';  
  }
}

async function cargarUsuarios() {
  const request = await fetch("http://localhost:5000/usuarios", {
    method: 'GET',
  })
  const usuarios = await request.json();
  console.log(usuarios)
  if (request.ok) {
            
            let listadoHtml = '';
            for (let usuario of usuarios) {
  
              let botonEliminar = '<a  href="#" onclick="eliminarUsuario(' + usuario.id + ')" class="btn btn-danger btn-circle btn-sm"><i class="fas fa-trash"></i></a>' ;
              let botonEditar = '<a  class="btn btn-primary" href="#" onclick="editarUsuario(' + usuario.id +')" btn  btn-circle btn-sm"><i class="bi bi-pencil-square"></i></a>  | ' ;
              let usuarioHtml = '<tr><td>'+ usuario.id+'</td><td>' + usuario.email + '</td><td>' + usuario.name + '</td><td>' + '</td><td>' + usuario.roles.id + '</td><td>'
                             +'</td><td>'+ botonEditar 
                               + botonEliminar + '</td></tr>';
              listadoHtml += usuarioHtml;
              }
            
            
            document.querySelector('#usuariosTBody').outerHTML = listadoHtml;
            
            }

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

const request = await fetch('http://localhost:5000/usuarios' + id, {
  method: 'DELETE',
  headers: getHeaders()
});
location.reload();
}
async function editarUsuario(id) {

  if (!confirm('¿Desea editar el usuario?')) {
    return;
  }
  
  const request = await fetch('http://localhost:5000/usuarios' + id, {
    method: 'POST',
    headers: getHeaders()
  });
  
}

async function asignaroRol(nombre) {

  if (!confirm('¿Desea asingar este rol al usuario?')) {
     return;
   }
 const request = await fetch('http://localhost:5000/usuarios/roles' + nombre, {
   method: 'POST',
   headers: getHeaders()
 });

 
 
 
} 

