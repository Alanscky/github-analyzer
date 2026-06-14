let chart = null;

async function searchUser() {

    const username = 
        document.getElementById("username").value;

    const response =
        await fetch(`/github/${username}`);

    if(!response.ok){
        alert("Usuario no encontrado");
        return;
    }

    const data =
        await response.json();

    console.log(data);

    document.getElementById("name")
        .textContent = data.name;

    document.getElementById("login")
        .textContent = data.login;

    document.getElementById("followers")
        .textContent = "Followers: " + data.followers;

    document.getElementById("repos")
        .textContent = "Repositories total: " + data.repos;

    document.getElementById("topLanguage")
        .textContent = "Main Language: " + data.language;

    document.getElementById("topRepo")
        .textContent = "Main repo: " + data.popularRepo;
    
    console.log(JSON.stringify(data.repositoryList, null, 2));

    renderChart(data.languages);


    renderRepositories(data.repositoryList);
}

function renderChart(languages) {

    const ctx =
        document.getElementById("languageChart");

    if (chart) {
        chart.destroy();
    }

    chart = new Chart(ctx, {
        type: "pie",

        data: {
            labels: Object.keys(languages),

            datasets: [{
                data: Object.values(languages)
            }]
        }
    });
}

function renderRepositories(repositories) {

    const list =
        document.getElementById("repositoryList");

    list.innerHTML = "";

    repositories.forEach(repo => {

        const li =
            document.createElement("li");

        li.textContent =
            `${repo.name} (${repo.language}): ${repo.description}`;

        list.appendChild(li);
    });
}

async function compareUsers() {

    const user1 =
        document.getElementById("compareUser1").value;

    const user2 =
        document.getElementById("compareUser2").value;

    const response =
        await fetch(
            `/github/compare/${user1}/${user2}`
        );

    if (!response.ok) {

        alert("Error al comparar usuarios");

        return;
    }

    const data =
        await response.json();

    renderComparison(data);
}

function renderComparison(data) {

    const result =
        document.getElementById(
            "comparisonResult"
        );

    result.innerHTML = `

        <h3>Comparación</h3>

        <table border="1">

            <tr>
                <th>Métrica</th>
                <th>${data.user1.login}</th>
                <th>${data.user2.login}</th>
            </tr>

            <tr>
                <td>Followers</td>
                <td>${data.user1.followers}</td>
                <td>${data.user2.followers}</td>
            </tr>

            <tr>
                <td>Repositorios</td>
                <td>${data.user1.repos}</td>
                <td>${data.user2.repos}</td>
            </tr>

            <tr>
                <td>Lenguaje Principal</td>
                <td>${data.user1.language}</td>
                <td>${data.user2.language}</td>
            </tr>

            <tr>
                <td>Repo Popular</td>
                <td>${data.user1.popularRepo}</td>
                <td>${data.user2.popularRepo}</td>
            </tr>

        </table>
    `;
}
