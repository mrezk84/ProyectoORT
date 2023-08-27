 // Call the dataTables jQuery plugin
$(document).ready(function() {
  asignarPermiso(id)
  cargarRoles()

  actualizarEmailDelUsuario();
});

function actualizarEmailDelUsuario() {
  document.getElementById('txt-email-usuario').outerHTML = request.nombre
}


async function cargarRoles() {
const request = await fetch('http://localhost:5000/roles', {
  method: 'GET',
  headers: getHeaders()
});
const roles = await request.json();


let listadoHtml = '';
for (let rol of roles) {
  let botonAsingar = '<a href="usuarios.html"  onclick =" asingarPermiso('+rol.id+')"  class="btn btn-primary"><span class="icon text-white-50"></i></span><span class="text">Asignar Rol</span></a>';
  let rolHtml = '<tr><td>'+ rol.id+'</td><td>' + rol.nombre +'</td><td>' + botonAsingar+ '</td></tr>';
  listadoHtml += rolHtml;
  }

document.querySelector('#roles tbody').outerHTML = listadoHtml;

}

function getHeaders() {
  return {
   'Accept': 'application/json',
   'Content-Type': 'application/json',
    
 };
}

 async function asignarPermiso(id) {

 if (!confirm('¿Desea asingar este rol al  usuario?')) {
    return;
  }
const request = await fetch('http://localhost:5000/usuarios/roles' + id, {
  method: 'POST',
  headers: getHeaders()
});
location.reload();

}
