$(document).ready(function() {
    getControles()
});


async function getControles() {
    const request = await fetch("http://localhost:5000/Controles/sinF", {
        method: 'GET',
    })

    const controles = await request.json();
    console.log(controles)
    if (request.ok) {
        let listadoHtml = '';
            for (let control of controles) {
                listadoHtml +=
                `
                 <option value="${control.id}">${control.id}--${formulario.descripcion}--${control.tipo}</option>
                `
        };

        document.getElementById("control").innerHTML = listadoHtml;

    }
}

async function agregarControl(){
    let formulario = null
    const url = new URL(document.URL);
    const searchParams = url.searchParams;
    formulario = searchParams.get('formulario_id');

    let control = document.getElementById('control').value;
    datos.formulario_id = Number(formulario);
    datos.control_id = Number(control);

    const request = await fetch("http://localhost:5000/controles/addForm", {
               method: 'POST',
               body: JSON.stringify(datos),
               headers: {
                   'Accept': 'application/json',
                   'Content-Type': 'application/json'
               },
           })
         if (request.status == 204) {
             alert("control agregado correctamente");
         }else{
             alert("Error agregando el control");
         }
             window.location.href = `controlesFormulario.html?formulario_id=${formulario}`;
}

// function redirectControlesFormulario(formulario){
//     window.location.href = `controlesFormulario.html?formulario=${encodeURIComponent(JSON.stringify(formulario))}`;
// }

// function redirectCrearControl(){
//     window.location.href = `altaControl.html`;
// }

// let formulario = null
// async function getControles() {
//     const request = await fetch('control', {
//         method: 'GET',
//     })

//     const url = new URL(document.URL);
//     const searchParams = url.searchParams;

//     formulario = searchParams.get('formulario');
//     console.log(formulario)
//     formulario = JSON.parse(decodeURIComponent(formulario));

//     let response = await request.json()
//     if (request.status == 200) {
//         let text = ``;
//         response.forEach(control => {
//             text +=
//                 `
//                  <tr>
//                     <th>${control.id}</th>
//                     <th>${control.descripcion}</th>
//                     <th><button onclick="agregarControl('${encodeURIComponent(JSON.stringify(control))}','${encodeURIComponent(JSON.stringify(formulario))}')">Agregar control</button></th>
//                 </tr>
//                 `
//         });
//         document.getElementById("controlesTBody").innerHTML = text;
//     }
// }

// async function agregarControl(control,formulario){
//     formulario = JSON.parse(decodeURIComponent(formulario));
//     control = JSON.parse(decodeURIComponent(control));
//     if(formulario.control.find(c => c.id == control.id)){
//         alert('El control ya se encuentra en el formulario')
//         return;
//     }
//     formulario.control.push(control);
//     const request = await fetch(`api/v1/formulario/${formulario.id}`, {
//         method: 'PUT',
//         headers: {
//             'Content-Type': 'application/json'
//         },
//         body: JSON.stringify(formulario)
//     })
//     if (request.status == 200) {
//         alert('Control agregado correctamente')
//         redirectControlesFormulario(formulario)
//     }
// }
