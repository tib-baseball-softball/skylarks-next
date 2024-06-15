export async function load({ parent }) {
    const data = await parent()

    return {
        clubTeams: data.clubTeams
    }
}