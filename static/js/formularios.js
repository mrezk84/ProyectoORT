

$(document).ready(function() {
  getFormularios()
});




async function getFormularios() {
  const request = await fetch(`api/v1/formulario`, {
    method: 'GET',
  })

  let response = await request.json()
  console.log(response)
  if (request.status == 200) {
            let text = ``
            response.forEach(formulario => {
                text +=
                    `
                     <tr>
                        <th>${formulario.nombre}</th>
                        <th><button onclick="redirectControles('${encodeURIComponent(JSON.stringify(formulario))}')">Editar Controles</button></th>
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
    window.location.href = `controlesFormulario.html?formulario=${formulario}`;
}