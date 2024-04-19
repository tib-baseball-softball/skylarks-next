export class LogoUtility {
    private static readonly FILE_PATH: string = '/TeamLogos/';

    private static readonly TEAM_LOGOS: { [key: string]: string } = {
        // our team comes first, so we always show up for SG instead of the other team
        'skylarks': LogoUtility.FILE_PATH + 'skylarks.svg',

        '89ers': LogoUtility.FILE_PATH + '89ers.png',
        'alligators': LogoUtility.FILE_PATH + 'alligators.png',
        'challengers': LogoUtility.FILE_PATH + 'challengers.png',
        'dockers': LogoUtility.FILE_PATH + 'dockers.png',
        'dragons': LogoUtility.FILE_PATH + 'dragons.png',
        'dukes': LogoUtility.FILE_PATH + 'dukes.png',
        'eagles': LogoUtility.FILE_PATH + 'eagles.png',
        'flamingos': LogoUtility.FILE_PATH + 'flamingos.png',
        'knights': LogoUtility.FILE_PATH + 'knights.png',
        'poorpigs': LogoUtility.FILE_PATH + 'poorpigs.png',
        'porcupines': LogoUtility.FILE_PATH + 'porcupines.png',
        'rams': LogoUtility.FILE_PATH + 'rams.png',
        'ravens': LogoUtility.FILE_PATH + 'ravens.png',
        'regents': LogoUtility.FILE_PATH + 'regents.png',
        'roadrunners': LogoUtility.FILE_PATH + 'roadrunners.png',
        'roosters': LogoUtility.FILE_PATH + 'roosters.png',
        'seahawks': LogoUtility.FILE_PATH + 'seahawks.svg',
        'sliders': LogoUtility.FILE_PATH + 'sliders.png',
        'sluggers': LogoUtility.FILE_PATH + 'sluggers.png',
        'stealers': LogoUtility.FILE_PATH + 'stealers.svg',
        'wild farmers': LogoUtility.FILE_PATH + 'wild_farmers.svg',
        'wizards': LogoUtility.FILE_PATH + 'wizards.png',
        'wallbreakers': LogoUtility.FILE_PATH + 'wallbreakers.png'
    };

    public static getLogoPathForTeamName(teamName: string): string {
        for (const key in LogoUtility.TEAM_LOGOS) {
            if (teamName.toLowerCase().includes(key)) {
                return LogoUtility.TEAM_LOGOS[key];
            }
        }
        return LogoUtility.FILE_PATH + 'default.svg';
    }
}
