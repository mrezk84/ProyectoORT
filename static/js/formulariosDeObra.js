

$(document).ready(function() {
  getFormularios()
});



let obra_id = null;
let piso_id  = null;
async function getFormularios() {
  const url = new URL(document.URL);
  const searchParams = url.searchParams;

  obra_id = searchParams.get('obra_id');
  piso_id = searchParams.get('piso_id');
  let cedula = localStorage.cedula;

  // todo añadir cedula/auth token al request
  // const request = await fetch(`check/formulario/${obra_id}/${piso_id}/${cedula}`, {
  const request = await fetch(`check/formulario/${obra_id}/${piso_id}`, {
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
                        <th><button onclick="redirectChecksByForm(${obra_id},${piso_id},${formulario.id})">Ver Controles</button></th>
                    </tr>
                    `
            });
            document.getElementById("formulariosTBody").innerHTML = text;

  }
}

function redirectAsignarFormularioAPiso(){
    window.location.href = `asignarFormularioAPiso.html?obra_id=${obra_id}&piso_id=${piso_id}`;
}



function redirectChecksByForm(obra_id,piso_id,formulario_id){
    window.location.href = `checkDeFormulario.html?obra_id=${obra_id}&piso_id=${piso_id}&formulario_id=${formulario_id}`;
}