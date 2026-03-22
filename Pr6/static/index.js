let counter = 0;
const scrollingElement = (document.scrollingElement || document.body);

document.getElementById('calculate').addEventListener('click', async () => {
    await getValuesToCalc();
});

document.getElementById('add').addEventListener('click', function () {
    let id = inputBlock(counter);

    let scrollDiv = document.getElementById(id).offsetTop;
    window.scrollTo({ top: scrollDiv, behavior: 'smooth'});

    counter++;
});

function inputBlock(counter){
    let blockId = 'user_input' + counter;
    let newInput = document.createElement("div");

    newInput.setAttribute('id', blockId);
    newInput.setAttribute('class', 'input_output');
    newInput.innerHTML = `
        <h3>Характеристики EП</h3>
        <div class="input_value">
            <label for="name`+ counter +`" class="left">Назва ЕП</label>
            <input id="name`+ counter +`" class="values name" type="text" required>
            <label for="name`+ counter +`" class="right"></label>
        </div>
        <div class="input_value">
            <label for="nominal_efficiency`+ counter +`" class="left"><abbr title="Номінальне значення коефіцієнта корисної дії ЕП">&eta;<sub>н</sub></abbr></label>
            <input id="nominal_efficiency`+ counter +`" class="values nominal_efficiency" type="number" step="0.01" required>
            <label for="nominal_efficiency`+ counter +`" class="right"></label>
        </div>
        <div class="input_value">
            <label for="load_cower_coef`+ counter +`" class="left"><abbr title="Коефіцієнт потужності навантаження">cos &phi;</abbr></label>
            <input id="load_cower_coef`+ counter +`" class="values load_cower_coef" type="number" step="0.01" required>
            <label for="load_cower_coef`+ counter +`" class="right"></label>
        </div>
        <div class="input_value">
            <label for="number`+ counter +`" class="left"><abbr title="Кількість ЕП">n</abbr></label>
            <input id="number`+ counter +`" class="values number" type="number" step="0.01"  required>
            <label for="number`+ counter +`" class="right">шт</label>
        </div>
        <div class="input_value">
            <label for="nominal_capacity`+ counter +`" class="left"><abbr title="Номінальна потужність ЕП:">P<sub>н</sub></abbr></label>
            <input id="nominal_capacity`+ counter +`" class="values nominal_capacity" type="number" step="0.01" required>
            <label for="nominal_capacity`+ counter +`" class="right">кВт</label>
        </div>
        <div class="input_value">
            <label for="utilization_factor`+ counter +`" class="left"><abbr title="Коефіцієнт використання">К<sub>В</sub></abbr></label>
            <input id="utilization_factor`+ counter +`" class="values utilization_factor" type="number" step="0.01" required>
            <label for="utilization_factor`+ counter +`" class="right"></label>
        </div>
        <div class="input_value">
            <label for="reactive_power_factor`+ counter +`" class="left"><abbr title="Коефіцієнт реактивної потужності">tg &phi;</abbr></label>
            <input id="reactive_power_factor`+ counter +`" class="values reactive_power_factor" type="number" step="0.01" required>
            <label for="reactive_power_factor`+ counter +`" class="right"></label>
        </div>
        <button id="delete" onclick="this.parentNode.parentNode.removeChild(this.parentNode)">X</button>`;
    document.getElementById('inputs').appendChild(newInput);
    return blockId;
}

function resultsBlock(result, loadVoltage) {
    const res = document.getElementById('results');
    if (!!res){
        res.parentNode.removeChild(res);
    }
    let newOutput = document.createElement("div");

    newOutput.setAttribute('id', 'results');
    newOutput.setAttribute('class', 'input_output');

    newOutput.innerHTML = `
        <div class="input_value">
            <p class="left">Груповий коефіцієнт використання для ШР1=ШР2=ШР3:</p>
            <p class="right">`+ result.GroupUtilizationFactor +`</p>
        </div>
        <div class="input_value">
            <p class="left">Ефективна кількість ЕП для ШР1=ШР2=ШР3:</p>
            <p class="right">`+ result.EffectiveNumber +`</p>
        </div>
        <div class="input_value">
            <p class="left">Розрахунковий коефіцієнт активної потужності для ШР1=ШР2=ШР3:</p>
            <p class="right">`+ result.EPCalculationFactor +`</p>
        </div>
        <div class="input_value">
            <p class="left">Розрахункове активне навантаження для ШР1=ШР2=ШР3:</p>
            <p class="right">`+ result.EstimatedActiveLoad +` кВт</p>
        </div>
        <div class="input_value">
            <p class="left">Розрахункове реактивне навантаження для ШР1=ШР2=ШР3:</p>
            <p class="right">`+ result.EstimatedReactiveLoad +` квар</p>
        </div>
        <div class="input_value">
            <p class="left">Повна потужність для ШР1=ШР2=ШР3:</p>
            <p class="right">`+ result.GroupFullPower +` кВ*А</p>
        </div>
        <div class="input_value">
            <p class="left">Розрахунковий груповий струм для ШР1=ШР2=ШР3:</p>
            <p class="right">`+ result.GroupCurrent +` А</p>
        </div>
        <div class="input_value">
            <p class="left">Коефіцієнти використання цеху в цілому:</p>
            <p class="right">`+ result.WorkshopUtilizationFactor +`</p>
        </div>
        <div class="input_value">
            <p class="left">Ефективна кількість ЕП цеху в цілому:</p>
            <p class="right">`+ result.WorkshopEffectiveNumber +`</p>
        </div>
        <div class="input_value">
            <p class="left">Розрахунковий коефіцієнт активної потужності цеху в цілому:</p>
            <p class="right">`+ result.WorkshopCalculationFactor +`</p>
        </div>
        <div class="input_value">
            <p class="left">Розрахункове активне навантаження на шинах `+ loadVoltage +` кВ ТП:</p>
            <p class="right">`+ result.EstimatedActiveTyreLoad +` кВт</p>
        </div>
        <div class="input_value">
            <p class="left">Розрахункове реактивне навантаження на шинах `+ loadVoltage +` кВ ТП:</p>
            <p class="right">`+ result.EstimatedReactiveTyreLoad +` квар</p>
        </div>
        <div class="input_value">
            <p class="left">Повна потужність на шинах `+ loadVoltage +` кВ ТП:</p>
            <p class="right">`+ result.WorkshopFullPower +` кВ*А</p>
        </div>
        <div class="input_value">
            <p class="left">Розрахунковий груповий струм на шинах `+ loadVoltage +` кВ ТП:</p>
            <p class="right">`+ result.WorkshopCurrent +` А</p>
        </div>`;
    document.querySelector('section').appendChild(newOutput);
}

async function getValuesToCalc(){

    let totalNumber = parseFloat(document.getElementById('number').value);
    if(Number.isNaN(totalNumber)){
        alert("Потрібно заповнити поле n цеху");
        return;
    }
    if(totalNumber < 0.0){
        alert("Було введене від'ємне значення n цеху");
        return;
    }

    let workshopNominalCapacity = parseFloat(document.getElementById('nominal_capacity').value);
    if(Number.isNaN(workshopNominalCapacity)){
        alert("Потрібно заповнити поле n*Pн");
        return;
    }
    if(workshopNominalCapacity < 0.0){
        alert("Було введене від'ємне значення n*Pн");
        return;
    }

    let workshopAverageActiveLoad = parseFloat(document.getElementById('average_active_load').value);
    if(Number.isNaN(workshopAverageActiveLoad)){
        alert("Потрібно заповнити поле n*Pн*Кв");
        return;
    }
    if(workshopAverageActiveLoad < 0.0){
        alert("Було введене від'ємне значення n*Pн*Кв");
        return;
    }

    let workshopAverageReactiveLoad = parseFloat(document.getElementById('average_reactive_load').value);
    if(Number.isNaN(workshopAverageReactiveLoad)){
        alert("Потрібно заповнити поле n*Pн*Кв*tg");
        return;
    }
    if(workshopAverageReactiveLoad < 0.0){
        alert("Було введене від'ємне значення n*Pн*Кв*tg");
        return;
    }

    let totalSquaredPower = parseFloat(document.getElementById('total_squared_power').value);
    if(Number.isNaN(totalSquaredPower)){
        alert("Потрібно заповнити поле n*Pн^2");
        return;
    }
    if(totalSquaredPower < 0.0){
        alert("Було введене від'ємне значення n*Pн^2");
        return;
    }

    let loadVoltage = parseFloat(document.getElementById('load_voltage').value);
    if(Number.isNaN(loadVoltage)){
        alert("Потрібно заповнити поле Uн");
        return;
    }
    if(loadVoltage < 0.0){
        alert("Було введене від'ємне значення Uн");
        return;
    }

    let name = [];
    let nominalEfficiency = [];
    let loadCowerFactor = [];
    let number = [];
    let nominalCapacity = [];
    let utilizationFactor = [];
    let reactivePowerFactor = [];
    for (let element of document.querySelectorAll('.name')) {
        let buf = element.value;
        if(buf === ''){
            alert("Потрібно заповнити ім'я");
            return;
        }
        name.push(buf);
    }
    for (let element of document.querySelectorAll('.nominal_efficiency')) {
        let buf = parseFloat(element.value);
        if(Number.isNaN(buf)){
            alert("Потрібно заповнити поле ККД");
            return;
        }
        if(buf < 0.0 || buf > 1.0){
            alert("Було введене недопустиме значення ККД");
            return;
        }
        nominalEfficiency.push(buf);
    }
    for (let element of document.querySelectorAll('.load_cower_coef')) {
        let buf = parseFloat(element.value);
        if(Number.isNaN(buf)){
            alert("Потрібно заповнити поле cos");
            return;
        }
        if(buf < 0.0 || buf > 1.0){
            alert("Було введене недопустиме значення cos");
            return;
        }
        loadCowerFactor.push(buf);
    }
    for (let element of document.querySelectorAll('.number')) {
        let buf = parseFloat(element.value);
        if(Number.isNaN(buf)){
            alert("Потрібно заповнити поле n");
            return;
        }
        if(buf < 0.0){
            alert("Було введене від'ємне значення n");
            return;
        }
        number.push(buf);
    }
    for (let element of document.querySelectorAll('.nominal_capacity')) {
        let buf = parseFloat(element.value);
        if(Number.isNaN(buf)){
            alert("Потрібно заповнити поле Pн");
            return;
        }
        if(buf < 0.0){
            alert("Було введене від'ємне значення Pн");
            return;
        }
        nominalCapacity.push(buf);
    }
    for (let element of document.querySelectorAll('.utilization_factor')) {
        let buf = parseFloat(element.value);
        if(Number.isNaN(buf)){
            alert("Потрібно заповнити поле Кв");
            return;
        }
        if(buf < 0.0){
            alert("Було введене від'ємне значення Кв");
            return;
        }
        utilizationFactor.push(buf);
    }
    for (let element of document.querySelectorAll('.reactive_power_factor')) {
        let buf = parseFloat(element.value);
        if(Number.isNaN(buf)){
            alert("Потрібно заповнити поле tg");
            return;
        }
        if(buf < 0.0){
            alert("Було введене від'ємне значення tg");
            return;
        }
        reactivePowerFactor.push(buf);
    }

    if (name.length === 0){
        alert('Треба додати хоч 1 ЕП');
        return;
    }

    const data = {
        "total_number": totalNumber,
        "total_nominal_capacity": workshopNominalCapacity,
        "average_active_load": workshopAverageActiveLoad,
        "average_reactive_load": workshopAverageReactiveLoad,
        "total_squared_power": totalSquaredPower,
        "load_voltage": loadVoltage,
        "name": name,
        "nominal_efficiency": nominalEfficiency,
        "load_cower_coef": loadCowerFactor,
        "number": number,
        "nominal_capacity": nominalCapacity,
        "utilization_factor": utilizationFactor,
        "reactive_power_factor": reactivePowerFactor
    }

    console.log(data)

    const result = await (await fetch("/calculate", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(data)
    })).json();

    resultsBlock(result, loadVoltage)
}
