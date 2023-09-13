$(document).ready(function() {
  getFormularios()
 
  $('#fromularios').DataTable();
});




async function getFormularios() {
  const request = await fetch("http://localhost:5000/formularios", {
    method: 'GET',
  })

  let response = await request.json()
  console.log(response)
  if (request.ok) {
            let text = ``
            response.forEach(formulario => {
                text +=
                    `
                     <tr>
                        <th>${formulario.nombre}</th>
                        <th>${formulario.informacion}</th>
                        <th>${formulario.version}</th>
                        <th>${formulario.fecha}</th>
                        <th><button onclick="redirectControles('${formulario.id}')" class="btn btn-primary btn-user"> Editar Controles </button>
                        <a onclick = "redirectUpdate('${formulario.id}')"  class="btn btn-success btn-icon-split"><span class="icon text-white-50"><i class="fas fa-check"></i></span>Update</a></th>
                    </tr>
                    `
            });
            document.getElementById("formulariosTBody").innerHTML = text;
  }
  
}

function redirectAltaFormulario(){
    window.location.href = `altaFormulario.html`;
}

async function redirectUpdate(id) {
  window.location.href = `formularioUpdate.html?formulario_id=${id}`;
}

function redirectControles(formulario){
    window.location.href = `controlesFormulario.html?formulario_id=${formulario}`;
}