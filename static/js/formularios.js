$(document).ready(function() {
  
  Paginar("fromularios");
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
                        <th><button onclick="redirectControles('${formulario.id}')" class="btn btn-primary btn-user"> Editar Controles </button></th>
                    </tr>
                    `
            });
            document.getElementById("formulariosTBody").innerHTML = text;
  }
  
}

function redirectAltaFormulario(){
    window.location.href = `altaFormulario.html`;
}

function redirectControles(formulario){
    window.location.href = `controlesFormulario.html?formulario_id=${formulario}`;
}
