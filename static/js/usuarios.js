// Call the dataTables jQuery plugin
$(document).ready(function() {
  cargarUsuarios();
  cargarRoles();
  actualizarCedulaDelUsuario();
$('#usuarios').DataTable();


});

function actualizarCedulaDelUsuario() {
  document.getElementById('txt-cedula-usuario').outerHTML = localStorage.cedula
}
async function cargarRoles() {
  const request = await fetch('api/roles', {
    method: 'GET',
    headers: getHeaders()
  });
  const roles = await request.json();
  if(roles.request==200){
   botonRol = '<a href="roles.html" onclick="asingarRol()" class="btn btn-success btn-icon-split"><span class="icon text-white-50"><i class="fas fa-check"></i></span><span class="text">Asignar</span></a>';  
  }
}

async function cargarUsuarios() {
const request = await fetch('api/usuarios', {
  method: 'GET',
  headers: getHeaders()
});
const usuarios = await request.json();


let listadoHtml = '';
for (let usuario of usuarios) {
  
  let botonEliminar = '<a href="#" onclick="eliminarUsuario(' + usuario.id + ')"  class="btn btn-danger btn-circle"><i class="fas fa-trash"></i></a>' ;
  let botonEditar = '<a href="#" onclick="editarUsuario(' + usuario.id +')" class="btn btn-info btn-circle"><i class="fas fa-info-circle"></i></a> |';  
 let botonAsignarRol=botonRol 
  let usuarioHtml = '<tr><td>'+ usuario.id+'</td><td>' + usuario.nombre + '</td><td>' + usuario.apellido + '</td><td>'
                  + usuario.cedula +'</td><td>'+ botonAsignarRol+'</td><td>'+ botonEditar 
                   + botonEliminar +  '</td></tr>';
  listadoHtml += usuarioHtml;
  }


document.querySelector('#usuarios tbody').outerHTML = listadoHtml;

}


function getHeaders() {
  return {
   'Accept': 'application/json',
   'Content-Type': 'application/json',
     'Authorization': localStorage.token  
 };
}

async function eliminarUsuario(id) {

if (!confirm('¿Desea eliminar este usuario?')) {
  return;
}

const request = await fetch('api/usuarios/' + id, {
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
  location.reload();
}

async function asignaroRol(nombre) {

  if (!confirm('¿Desea asingar este rol al usuario?')) {
     return;
   }
 const request = await fetch('api/roles' + nombre, {
   method: 'POST',
   headers: getHeaders()
 });

 
 
 location.reload();
} 

