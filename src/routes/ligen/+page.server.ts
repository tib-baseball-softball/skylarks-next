export async function load({ parent }) {
    const data = await parent()

    return {
        leagueGroups: data.leagueGroups
    }
}