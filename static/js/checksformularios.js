

$(document).ready(function() {
  getCheckFormularios()
});


async function getCheckFormularios() {
  const url = new URL(document.URL);
  const searchParams = url.searchParams;

  let obra_id = searchParams.get('obra_id');
  let piso_id = searchParams.get('piso_id');
  let form_id = searchParams.get('formulario_id');

  const request = await fetch(`check/${form_id}/${obra_id}/${piso_id}`, {
    method: 'GET',
  })

  let response = await request.json()
  console.log(response)
  if (request.status == 200) {
            let text = ``
            response.forEach(check => {
                text +=
                    `
                     <tr>
                        <th>${check.id}</th>
                        <th>${check.estado == null ? "Pendiente":check.estado}</th>
                        <th>${check.fechaControl == null ? "":check.fechaControl}</th>
                        <th>${check.observaciones == null ? "":check.observaciones}</th>
                        <th>${check.version == null ? "":check.version}</th>
                        <th><button onclick="setCheckCompletado(${check.id},'${encodeURIComponent(JSON.stringify(response))}')">Completado</button></th>
                        <th><button onclick="setCheckNoCompletado(${check.id},'${encodeURIComponent(JSON.stringify(response))}')">No completado</button></th>
                    </tr>
                    `
            });
            document.getElementById("checksformulariosTBody").innerHTML = text;
  }
}



async function setCheckCompletado(check_id,response){
    response = JSON.parse(decodeURIComponent(response));
    const request = await fetch(`check/${check_id}?estado=COMPLETADO&fecha=${getNowDate()}`, {
        method: 'PUT',
    })
    let text =``;
    response.forEach(check => {
        if (check.id == check_id){
            check.estado = "COMPLETADO";
            check.fechaControl = getNowDate();
        }
        text +=
            `
                     <tr>
                        <th>${check.id}</th>
                        <th>${check.estado == null ? "Pendiente":check.estado}</th>
                        <th>${check.fechaControl == null ? "":check.fechaControl}</th>
                        <th>${check.observaciones == null ? "":check.observaciones}</th>
                        <th>${check.version == null ? "":check.version}</th>
                        <th><button onclick="setCheckCompletado(${check.id},'${encodeURIComponent(JSON.stringify(response))}')">Completado</button></th>
                        <th><button onclick="setCheckNoCompletado(${check.id},'${encodeURIComponent(JSON.stringify(response))}')">No completado</button></th>
                    </tr>
                    `
    });
    document.getElementById("checksformulariosTBody").innerHTML = text;

}
async function setCheckNoCompletado(check_id,response){
    response = JSON.parse(decodeURIComponent(response));
    const request = await fetch(`check/${check_id}?estado=NO_COMPLETADO&fecha=${getNowDate()}`, {
        method: 'PUT',
    })
    let text =``;
    response.forEach(check => {
        if (check.id == check_id){
            check.estado = "NO_COMPLETADO";
            check.fechaControl = getNowDate();
        }
        text +=
            `
                     <tr>
                        <th>${check.id}</th>
                        <th>${check.estado == null ? "Pendiente":check.estado}</th>
                        <th>${check.fechaControl == null ? "":check.fechaControl}</th>
                        <th>${check.observaciones == null ? "":check.observaciones}</th>
                        <th>${check.version == null ? "":check.version}</th>
                        <th><button onclick="setCheckCompletado(${check.id},'${encodeURIComponent(JSON.stringify(response))}')">Completado</button></th>
                        <th><button onclick="setCheckNoCompletado(${check.id},'${encodeURIComponent(JSON.stringify(response))}')">No completado</button></th>
                    </tr>
                    `
    });
    document.getElementById("checksformulariosTBody").innerHTML = text;
}

function getNowDate(){
    return new Date(Date.now()).toISOString().split('T')[0]
}