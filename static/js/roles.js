 // Call the dataTables jQuery plugin
$(document).ready(function() {
  cargarRoles();
 
$('#roles').DataTable();

actualizarCedulaDelUsuario();
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


let listadoHtml = '';
for (let rol of roles) {
  let botonAsingar = '<a href="usuarios.html"  onclick =" asingarPermiso('+rol.id+')"  class="btn btn-success btn-icon-split"><span class="icon text-white-50"><i class="fas fa-check"></i></span><span class="text">Asignar Rol</span></a>';
  let rolHtml = '<tr><td>'+ rol.id+'</td><td>' + rol.nombreRol +'</td><td>' + botonAsingar+ '</td></tr>';
  listadoHtml += rolHtml;
  }

document.querySelector('#roles tbody').outerHTML = listadoHtml;

}

function getHeaders() {
  return {
   'Accept': 'application/json',
   'Content-Type': 'application/json',
     'Authorization': localStorage.token  
 };
}

 async function asignarPermiso(id) {

 if (!confirm('¿Desea asingar este rol al  usuario?')) {
    return;
  }
const request = await fetch('api/roles' + id, {
  method: 'POST',
  headers: getHeaders()
});
location.reload();

}
