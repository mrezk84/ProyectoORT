

$(document).ready(function() {
  getFormularios()
});


let obra_id = null;
let piso_id  = null;

async function getFormularios() {
  const request = await fetch(`api/v1/formulario`, {
    method: 'GET',
  })
    const url = new URL(document.URL);
    const searchParams = url.searchParams;

    obra_id = searchParams.get('obra_id');
    piso_id = searchParams.get('piso_id');

    let response = await request.json()
  if (request.status == 200) {
            let text = ``
            response.forEach(formulario => {
                text +=
                    `
                     <tr>
                        <th>${formulario.nombre}</th>
                        <th><button onclick="asignarFormulario('${encodeURIComponent(JSON.stringify(formulario))}')">Asignar formulario</button></th>
                    </tr>
                    `
            });
            document.getElementById("formulariosTBody").innerHTML = text;
  }
}

async function asignarFormulario(formulario){
    formulario = JSON.parse(decodeURIComponent(formulario));
    const request = await fetch(`check/assign/${formulario.id}/${obra_id}/${piso_id}`, {
        method: 'POST',
    });
    if (request.status == 200) {
        alert("Formulario asignado correctamente");
    }else {
        alert("Error al asignar formulario");
    }
        window.location.href = `formulariosDeObra.html?obra_id=${obra_id}&piso_id=${piso_id}`;
}