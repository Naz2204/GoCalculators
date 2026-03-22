document.getElementById('fuel_list').addEventListener('input', onChangeFuelType);
document.addEventListener("DOMContentLoaded", onLoadFuelType)

function onChangeFuelType() {
    if (this.value === 'user') {
        changeRequired(true)
        document.getElementById('low_burn_temp').value = '';
        document.getElementById('ash_part').value = '';
        document.getElementById('ash_mass').value = '';
        document.getElementById('burnable_ejection').value = '';
        document.getElementById('cleaning_efficiency').value = '';
        document.getElementById('solid_emission_with_sulfur').value = '';
        document.getElementById('mass').value = '';
        document.getElementById('user_input').style.display = 'flex';
        document.getElementById('emission').innerHTML = '';
        document.getElementById('ejection').innerHTML = '';
    }
    else {
        document.getElementById('mass').value = '';
        changeRequired(false);
        document.getElementById('user_input').style.display = 'none';
        document.getElementById('emission').innerHTML = '';
        document.getElementById('ejection').innerHTML = '';
    }
}

function onLoadFuelType() {
    if (document.getElementById('fuel_list').value === 'user') {
        changeRequired(true)
        document.getElementById('user_input').style.display = 'flex';
    }
    else {
        changeRequired(false);
        document.getElementById('user_input').style.display = 'none';
    }
}


function changeRequired(state) {
    document.getElementById('low_burn_temp').required = state;
    document.getElementById('ash_part').required = state;
    document.getElementById('ash_mass').required = state;
    document.getElementById('burnable_ejection').required = state;
    document.getElementById('cleaning_efficiency').required = state;
    document.getElementById('solid_emission_with_sulfur').required = state;
}