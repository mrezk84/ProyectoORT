$(document).ready(function() {
    getControles()
});

let formulario = null
async function getControles() {
    const request = await fetch('control', {
        method: 'GET',
    })

    const url = new URL(document.URL);
    const searchParams = url.searchParams;

    formulario = searchParams.get('formulario');
    console.log(formulario)
    formulario = JSON.parse(decodeURIComponent(formulario));

    let response = await request.json()
    if (request.status == 200) {
        let text = ``;
        response.forEach(control => {
            text +=
                `
                 <tr>
                    <th>${control.id}</th>
                    <th>${control.descripcion}</th>
                    <th><button onclick="agregarControl('${encodeURIComponent(JSON.stringify(control))}','${encodeURIComponent(JSON.stringify(formulario))}')">Agregar control</button></th>
                </tr>
                `
        });
        document.getElementById("controlesTBody").innerHTML = text;
    }
}

async function agregarControl(control,formulario){
    formulario = JSON.parse(decodeURIComponent(formulario));
    control = JSON.parse(decodeURIComponent(control));
    if(formulario.control.find(c => c.id == control.id)){
        alert('El control ya se encuentra en el formulario')
        return;
    }
    formulario.control.push(control);
    const request = await fetch(`api/v1/formulario/${formulario.id}`, {
        method: 'PUT',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(formulario)
    })
    if (request.status == 200) {
        alert('Control agregado correctamente')
        redirectControlesFormulario(formulario)
    }
}
function redirectControlesFormulario(formulario){
    window.location.href = `controlesFormulario.html?formulario=${encodeURIComponent(JSON.stringify(formulario))}`;
}
function redirectCrearControl(){
    window.location.href = `altaControl.html`;
}