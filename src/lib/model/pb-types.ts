// === start of custom type ===
  // Events.EventsMatch_json.match_json
  export type EventsMatch_json = Array<{
 
  }>;
  // === end of custom type ===


// === start of custom type ===
  // Requestcaches.RequestcachesResponseBody.responseBody
  export type RequestcachesResponseBody = Array<{
 
  }>;
  // === end of custom type ===


/**
 * This file was @generated using typed-pocketbase
 */

// https://pocketbase.io/docs/collections/#base-collection
export interface BaseCollectionResponse {
	/**
	 * 15 characters string to store as record ID.
	 */
	id: string;
	/**
	 * Date string representation for the creation date.
	 */
	created: string;
	/**
	 * Date string representation for the creation date.
	 */
	updated: string;
	/**
	 * The collection id.
	 */
	collectionId: string;
	/**
	 * The collection name.
	 */
	collectionName: string;
}

// https://pocketbase.io/docs/api-records/#create-record
export interface BaseCollectionCreate {
	/**
	 * 15 characters string to store as record ID.
	 * If not set, it will be auto generated.
	 */
	id?: string;
}

// https://pocketbase.io/docs/api-records/#update-record
export interface BaseCollectionUpdate {}

// https://pocketbase.io/docs/collections/#auth-collection
export interface AuthCollectionResponse extends BaseCollectionResponse {
	/**
	 * The username of the auth record.
	 */
	username: string;
	/**
	 * Auth record email address.
	 */
	email: string;
	/**
	 * Auth record email address.
	 */
	tokenKey?: string;
	/**
	 * Whether to show/hide the auth record email when fetching the record data.
	 */
	emailVisibility: boolean;
	/**
	 * Indicates whether the auth record is verified or not.
	 */
	verified: boolean;
}

// https://pocketbase.io/docs/api-records/#create-record
export interface AuthCollectionCreate extends BaseCollectionCreate {
	/**
	 * The username of the auth record.
	 * If not set, it will be auto generated.
	 */
	username?: string;
	/**
	 * Auth record email address.
	 */
	email?: string;
	/**
	 * Whether to show/hide the auth record email when fetching the record data.
	 */
	emailVisibility?: boolean;
	/**
	 * Auth record password.
	 */
	password: string;
	/**
	 * Auth record password confirmation.
	 */
	passwordConfirm: string;
	/**
	 * Indicates whether the auth record is verified or not.
	 * This field can be set only by admins or auth records with "Manage" access.
	 */
	verified?: boolean;
}

// https://pocketbase.io/docs/api-records/#update-record
export interface AuthCollectionUpdate {
	/**
	 * The username of the auth record.
	 */
	username?: string;
	/**
	 * The auth record email address.
	 * This field can be updated only by admins or auth records with "Manage" access.
	 * Regular accounts can update their email by calling "Request email change".
	 */
	email?: string;
	/**
	 * Whether to show/hide the auth record email when fetching the record data.
	 */
	emailVisibility?: boolean;
	/**
	 * Old auth record password.
	 * This field is required only when changing the record password. Admins and auth records with "Manage" access can skip this field.
	 */
	oldPassword?: string;
	/**
	 * New auth record password.
	 */
	password?: string;
	/**
	 * New auth record password confirmation.
	 */
	passwordConfirm?: string;
	/**
	 * Indicates whether the auth record is verified or not.
	 * This field can be set only by admins or auth records with "Manage" access.
	 */
	verified?: boolean;
}

// https://pocketbase.io/docs/collections/#view-collection
export interface ViewCollectionRecord {
	id: string;
}

// utilities

type MaybeArray<T> = T | T[];
// ==== start of users block =====


export interface UsersResponse extends AuthCollectionResponse {
	collectionName: 'users';
	id: string;
	tokenKey: string;
	email: string;
	emailVisibility: boolean;
	verified: boolean;
	username: string;
	first_name: string;
	last_name: string;
	avatar: string;
	teams: Array<string>;
	club: Array<string>;
	last_login: string;
	birthday: string;
	number: string;
	position: Array<'0' | '1' | '2' | '3' | '4' | '5' | '6' | '7' | '8' | '9' | '10' | '11' | '12' | '13' | '14'>;
	bats: '' | '0' | '1' | '2' | '3';
	throws: '' | '0' | '1' | '2' | '3';
	image: string;
	umpire: '' | '0' | '1' | '2' | '3' | '4';
	scorer: '' | '0' | '1' | '2' | '3';
	bsm_id: number;
	signup_key: string;
	created: string;
	updated: string;
}

export interface UsersCreate extends AuthCollectionCreate {
	id: string;
	password: string;
	tokenKey: string;
	email?: string;
	emailVisibility?: boolean;
	verified?: boolean;
	username: string;
	first_name?: string;
	last_name?: string;
	avatar?: File | null;
	teams?: MaybeArray<string>;
	club?: MaybeArray<string>;
	last_login?: string | Date;
	birthday?: string | Date;
	number?: string;
	position?: MaybeArray<'0' | '1' | '2' | '3' | '4' | '5' | '6' | '7' | '8' | '9' | '10' | '11' | '12' | '13' | '14'>;
	bats?: '' | '0' | '1' | '2' | '3';
	throws?: '' | '0' | '1' | '2' | '3';
	image?: File | null;
	umpire?: '' | '0' | '1' | '2' | '3' | '4';
	scorer?: '' | '0' | '1' | '2' | '3';
	bsm_id?: number;
	signup_key?: string;
	created?: string | Date;
	updated?: string | Date;
}

export interface UsersUpdate extends AuthCollectionUpdate {
	id?: string;
	tokenKey?: string;
	email?: string;
	emailVisibility?: boolean;
	verified?: boolean;
	username?: string;
	first_name?: string;
	last_name?: string;
	avatar?: File | null;
	teams?: MaybeArray<string>;
	'teams+'?: MaybeArray<string>;
	'teams-'?: MaybeArray<string>;
	club?: MaybeArray<string>;
	'club+'?: MaybeArray<string>;
	'club-'?: MaybeArray<string>;
	last_login?: string | Date;
	birthday?: string | Date;
	number?: string;
	position?: MaybeArray<'0' | '1' | '2' | '3' | '4' | '5' | '6' | '7' | '8' | '9' | '10' | '11' | '12' | '13' | '14'>;
	'position+'?: MaybeArray<'0' | '1' | '2' | '3' | '4' | '5' | '6' | '7' | '8' | '9' | '10' | '11' | '12' | '13' | '14'>;
	'position-'?: MaybeArray<'0' | '1' | '2' | '3' | '4' | '5' | '6' | '7' | '8' | '9' | '10' | '11' | '12' | '13' | '14'>;
	bats?: '' | '0' | '1' | '2' | '3';
	throws?: '' | '0' | '1' | '2' | '3';
	image?: File | null;
	umpire?: '' | '0' | '1' | '2' | '3' | '4';
	scorer?: '' | '0' | '1' | '2' | '3';
	bsm_id?: number;
	'bsm_id+'?: number;
	'bsm_id-'?: number;
	signup_key?: string;
	created?: string | Date;
	updated?: string | Date;
}

export interface UsersCollection {
	type: 'auth';
	collectionId: string;
	collectionName: 'users';
	response: UsersResponse;
	create: UsersCreate;
	update: UsersUpdate;
	relations: {
		teams: TeamsCollection[];
		club: ClubsCollection[];
		clubs_via_admins: ClubsCollection[];
		teams_via_admins: TeamsCollection[];
		participations_via_user: ParticipationsCollection[];
	};
}

// ==== end of users block =====

// ==== start of events block =====


export interface EventsResponse extends BaseCollectionResponse {
	collectionName: 'events';
	id: string;
	bsm_id: number;
	title: string;
	desc: string;
	starttime: string;
	meetingtime: string;
	endtime: string;
	type: 'game' | 'practice' | 'misc';
	attire: string;
	team: string;
	cancelled: boolean;
	match_json?: EventsMatch_json
	guests: string;
	location: string;
	series: string;
	created: string;
	updated: string;
}

export interface EventsCreate extends BaseCollectionCreate {
	id: string;
	bsm_id?: number;
	title: string;
	desc?: string;
	starttime: string | Date;
	meetingtime: string | Date;
	endtime?: string | Date;
	type: 'game' | 'practice' | 'misc';
	attire?: string;
	team: string;
	cancelled?: boolean;
	match_json?: Record<string, any> | Array<any> | null;
	guests?: string;
	location?: string;
	series?: string;
	created?: string | Date;
	updated?: string | Date;
}

export interface EventsUpdate extends BaseCollectionUpdate {
	id?: string;
	bsm_id?: number;
	'bsm_id+'?: number;
	'bsm_id-'?: number;
	title?: string;
	desc?: string;
	starttime?: string | Date;
	meetingtime?: string | Date;
	endtime?: string | Date;
	type?: 'game' | 'practice' | 'misc';
	attire?: string;
	team?: string;
	cancelled?: boolean;
	match_json?: Record<string, any> | Array<any> | null;
	guests?: string;
	location?: string;
	series?: string;
	created?: string | Date;
	updated?: string | Date;
}

export interface EventsCollection {
	type: 'base';
	collectionId: string;
	collectionName: 'events';
	response: EventsResponse;
	create: EventsCreate;
	update: EventsUpdate;
	relations: {
		attire: UniformsetsCollection;
		team: TeamsCollection;
		series: EventseriesCollection;
		participations_via_event: ParticipationsCollection[];
	};
}

// ==== end of events block =====

// ==== start of clubs block =====


export interface ClubsResponse extends BaseCollectionResponse {
	collectionName: 'clubs';
	id: string;
	name: string;
	bsm_id: number;
	bsm_api_key: string;
	signup_key: string;
	acronym: string;
	admins: Array<string>;
	created: string;
	updated: string;
}

export interface ClubsCreate extends BaseCollectionCreate {
	id: string;
	name: string;
	bsm_id: number;
	bsm_api_key?: string;
	signup_key: string;
	acronym?: string;
	admins: MaybeArray<string>;
	created?: string | Date;
	updated?: string | Date;
}

export interface ClubsUpdate extends BaseCollectionUpdate {
	id?: string;
	name?: string;
	bsm_id?: number;
	'bsm_id+'?: number;
	'bsm_id-'?: number;
	bsm_api_key?: string;
	signup_key?: string;
	acronym?: string;
	admins?: MaybeArray<string>;
	'admins+'?: MaybeArray<string>;
	'admins-'?: MaybeArray<string>;
	created?: string | Date;
	updated?: string | Date;
}

export interface ClubsCollection {
	type: 'base';
	collectionId: string;
	collectionName: 'clubs';
	response: ClubsResponse;
	create: ClubsCreate;
	update: ClubsUpdate;
	relations: {
		users_via_club: UsersCollection[];
		admins: UsersCollection[];
		teams_via_club: TeamsCollection[];
		leaguegroups_via_clubs: LeaguegroupsCollection[];
		uniformsets_via_club: UniformsetsCollection[];
	};
}

// ==== end of clubs block =====

// ==== start of teams block =====


export interface TeamsResponse extends BaseCollectionResponse {
	collectionName: 'teams';
	id: string;
	name: string;
	bsm_league_group: number;
	age_group: 'adults' | 'minors';
	club: string;
	description: string;
	admins: Array<string>;
	created: string;
	updated: string;
}

export interface TeamsCreate extends BaseCollectionCreate {
	id: string;
	name: string;
	bsm_league_group?: number;
	age_group: 'adults' | 'minors';
	club: string;
	description?: string;
	admins?: MaybeArray<string>;
	created?: string | Date;
	updated?: string | Date;
}

export interface TeamsUpdate extends BaseCollectionUpdate {
	id?: string;
	name?: string;
	bsm_league_group?: number;
	'bsm_league_group+'?: number;
	'bsm_league_group-'?: number;
	age_group?: 'adults' | 'minors';
	club?: string;
	description?: string;
	admins?: MaybeArray<string>;
	'admins+'?: MaybeArray<string>;
	'admins-'?: MaybeArray<string>;
	created?: string | Date;
	updated?: string | Date;
}

export interface TeamsCollection {
	type: 'base';
	collectionId: string;
	collectionName: 'teams';
	response: TeamsResponse;
	create: TeamsCreate;
	update: TeamsUpdate;
	relations: {
		users_via_teams: UsersCollection[];
		events_via_team: EventsCollection[];
		club: ClubsCollection;
		admins: UsersCollection[];
		eventseries_via_team: EventseriesCollection[];
	};
}

// ==== end of teams block =====

// ==== start of requestcaches block =====


export interface RequestcachesResponse extends BaseCollectionResponse {
	collectionName: 'requestcaches';
	id: string;
	hash: string;
	responseBody?: RequestcachesResponseBody
	url: string;
	created: string;
	updated: string;
}

export interface RequestcachesCreate extends BaseCollectionCreate {
	id: string;
	hash: string;
	responseBody?: RequestcachesResponseBody
	url: string | URL;
	created?: string | Date;
	updated?: string | Date;
}

export interface RequestcachesUpdate extends BaseCollectionUpdate {
	id?: string;
	hash?: string;
	responseBody?: RequestcachesResponseBody
	url?: string | URL;
	created?: string | Date;
	updated?: string | Date;
}

export interface RequestcachesCollection {
	type: 'base';
	collectionId: string;
	collectionName: 'requestcaches';
	response: RequestcachesResponse;
	create: RequestcachesCreate;
	update: RequestcachesUpdate;
	relations: Record<string, never>;
}

// ==== end of requestcaches block =====

// ==== start of leaguegroups block =====


export interface LeaguegroupsResponse extends BaseCollectionResponse {
	collectionName: 'leaguegroups';
	id: string;
	bsm_id: number;
	season: number;
	acronym: string;
	name: string;
	clubs: Array<string>;
	created: string;
	updated: string;
}

export interface LeaguegroupsCreate extends BaseCollectionCreate {
	id: string;
	bsm_id: number;
	season: number;
	acronym: string;
	name: string;
	clubs?: MaybeArray<string>;
	created?: string | Date;
	updated?: string | Date;
}

export interface LeaguegroupsUpdate extends BaseCollectionUpdate {
	id?: string;
	bsm_id?: number;
	'bsm_id+'?: number;
	'bsm_id-'?: number;
	season?: number;
	'season+'?: number;
	'season-'?: number;
	acronym?: string;
	name?: string;
	clubs?: MaybeArray<string>;
	'clubs+'?: MaybeArray<string>;
	'clubs-'?: MaybeArray<string>;
	created?: string | Date;
	updated?: string | Date;
}

export interface LeaguegroupsCollection {
	type: 'base';
	collectionId: string;
	collectionName: 'leaguegroups';
	response: LeaguegroupsResponse;
	create: LeaguegroupsCreate;
	update: LeaguegroupsUpdate;
	relations: {
		clubs: ClubsCollection[];
	};
}

// ==== end of leaguegroups block =====

// ==== start of participations block =====


export interface ParticipationsResponse extends BaseCollectionResponse {
	collectionName: 'participations';
	id: string;
	user: string;
	event: string;
	comment: string;
	state: '' | 'in' | 'out' | 'maybe';
	created: string;
	updated: string;
}

export interface ParticipationsCreate extends BaseCollectionCreate {
	id: string;
	user?: string;
	event?: string;
	comment?: string;
	state?: '' | 'in' | 'out' | 'maybe';
	created?: string | Date;
	updated?: string | Date;
}

export interface ParticipationsUpdate extends BaseCollectionUpdate {
	id?: string;
	user?: string;
	event?: string;
	comment?: string;
	state?: '' | 'in' | 'out' | 'maybe';
	created?: string | Date;
	updated?: string | Date;
}

export interface ParticipationsCollection {
	type: 'base';
	collectionId: string;
	collectionName: 'participations';
	response: ParticipationsResponse;
	create: ParticipationsCreate;
	update: ParticipationsUpdate;
	relations: {
		user: UsersCollection;
		event: EventsCollection;
	};
}

// ==== end of participations block =====

// ==== start of uniformsets block =====


export interface UniformsetsResponse extends BaseCollectionResponse {
	collectionName: 'uniformsets';
	id: string;
	name: string;
	cap: string;
	jersey: string;
	pants: string;
	club: string;
	created: string;
	updated: string;
}

export interface UniformsetsCreate extends BaseCollectionCreate {
	id: string;
	name: string;
	cap: string;
	jersey: string;
	pants: string;
	club: string;
	created?: string | Date;
	updated?: string | Date;
}

export interface UniformsetsUpdate extends BaseCollectionUpdate {
	id?: string;
	name?: string;
	cap?: string;
	jersey?: string;
	pants?: string;
	club?: string;
	created?: string | Date;
	updated?: string | Date;
}

export interface UniformsetsCollection {
	type: 'base';
	collectionId: string;
	collectionName: 'uniformsets';
	response: UniformsetsResponse;
	create: UniformsetsCreate;
	update: UniformsetsUpdate;
	relations: {
		events_via_attire: EventsCollection[];
		club: ClubsCollection;
	};
}

// ==== end of uniformsets block =====

// ==== start of _superusers block =====


export interface SuperusersResponse extends AuthCollectionResponse {
	collectionName: '_superusers';
	id: string;
	tokenKey: string;
	email: string;
	emailVisibility: boolean;
	verified: boolean;
	created: string;
	updated: string;
}

export interface SuperusersCreate extends AuthCollectionCreate {
	id: string;
	password: string;
	tokenKey: string;
	email: string;
	emailVisibility?: boolean;
	verified?: boolean;
	created?: string | Date;
	updated?: string | Date;
}

export interface SuperusersUpdate extends AuthCollectionUpdate {
	id?: string;
	tokenKey?: string;
	email?: string;
	emailVisibility?: boolean;
	verified?: boolean;
	created?: string | Date;
	updated?: string | Date;
}

export interface SuperusersCollection {
	type: 'auth';
	collectionId: string;
	collectionName: '_superusers';
	response: SuperusersResponse;
	create: SuperusersCreate;
	update: SuperusersUpdate;
	relations: Record<string, never>;
}

// ==== end of _superusers block =====

// ==== start of _externalAuths block =====


export interface ExternalAuthsResponse extends BaseCollectionResponse {
	collectionName: '_externalAuths';
	id: string;
	collectionRef: string;
	recordRef: string;
	provider: string;
	providerId: string;
	created: string;
	updated: string;
}

export interface ExternalAuthsCreate extends BaseCollectionCreate {
	id: string;
	collectionRef: string;
	recordRef: string;
	provider: string;
	providerId: string;
	created?: string | Date;
	updated?: string | Date;
}

export interface ExternalAuthsUpdate extends BaseCollectionUpdate {
	id?: string;
	collectionRef?: string;
	recordRef?: string;
	provider?: string;
	providerId?: string;
	created?: string | Date;
	updated?: string | Date;
}

export interface ExternalAuthsCollection {
	type: 'base';
	collectionId: string;
	collectionName: '_externalAuths';
	response: ExternalAuthsResponse;
	create: ExternalAuthsCreate;
	update: ExternalAuthsUpdate;
	relations: Record<string, never>;
}

// ==== end of _externalAuths block =====

// ==== start of _mfas block =====


export interface MfasResponse extends BaseCollectionResponse {
	collectionName: '_mfas';
	id: string;
	collectionRef: string;
	recordRef: string;
	method: string;
	created: string;
	updated: string;
}

export interface MfasCreate extends BaseCollectionCreate {
	id: string;
	collectionRef: string;
	recordRef: string;
	method: string;
	created?: string | Date;
	updated?: string | Date;
}

export interface MfasUpdate extends BaseCollectionUpdate {
	id?: string;
	collectionRef?: string;
	recordRef?: string;
	method?: string;
	created?: string | Date;
	updated?: string | Date;
}

export interface MfasCollection {
	type: 'base';
	collectionId: string;
	collectionName: '_mfas';
	response: MfasResponse;
	create: MfasCreate;
	update: MfasUpdate;
	relations: Record<string, never>;
}

// ==== end of _mfas block =====

// ==== start of _otps block =====


export interface OtpsResponse extends BaseCollectionResponse {
	collectionName: '_otps';
	id: string;
	collectionRef: string;
	recordRef: string;
	created: string;
	updated: string;
	sentTo: string;
}

export interface OtpsCreate extends BaseCollectionCreate {
	id: string;
	collectionRef: string;
	recordRef: string;
	password: string;
	created?: string | Date;
	updated?: string | Date;
	sentTo?: string;
}

export interface OtpsUpdate extends BaseCollectionUpdate {
	id?: string;
	collectionRef?: string;
	recordRef?: string;
	created?: string | Date;
	updated?: string | Date;
	sentTo?: string;
}

export interface OtpsCollection {
	type: 'base';
	collectionId: string;
	collectionName: '_otps';
	response: OtpsResponse;
	create: OtpsCreate;
	update: OtpsUpdate;
	relations: Record<string, never>;
}

// ==== end of _otps block =====

// ==== start of _authOrigins block =====


export interface AuthOriginsResponse extends BaseCollectionResponse {
	collectionName: '_authOrigins';
	id: string;
	collectionRef: string;
	recordRef: string;
	fingerprint: string;
	created: string;
	updated: string;
}

export interface AuthOriginsCreate extends BaseCollectionCreate {
	id: string;
	collectionRef: string;
	recordRef: string;
	fingerprint: string;
	created?: string | Date;
	updated?: string | Date;
}

export interface AuthOriginsUpdate extends BaseCollectionUpdate {
	id?: string;
	collectionRef?: string;
	recordRef?: string;
	fingerprint?: string;
	created?: string | Date;
	updated?: string | Date;
}

export interface AuthOriginsCollection {
	type: 'base';
	collectionId: string;
	collectionName: '_authOrigins';
	response: AuthOriginsResponse;
	create: AuthOriginsCreate;
	update: AuthOriginsUpdate;
	relations: Record<string, never>;
}

// ==== end of _authOrigins block =====

// ==== start of eventseries block =====


export interface EventseriesResponse extends BaseCollectionResponse {
	collectionName: 'eventseries';
	id: string;
	interval: number;
	weekday: number;
	end: string;
	team: string;
	start: string;
	time: string;
	created: string;
	updated: string;
}

export interface EventseriesCreate extends BaseCollectionCreate {
	id: string;
	interval: number;
	weekday?: number;
	end?: string | Date;
	team: string;
	start: string | Date;
	time: string;
	created?: string | Date;
	updated?: string | Date;
}

export interface EventseriesUpdate extends BaseCollectionUpdate {
	id?: string;
	interval?: number;
	'interval+'?: number;
	'interval-'?: number;
	weekday?: number;
	'weekday+'?: number;
	'weekday-'?: number;
	end?: string | Date;
	team?: string;
	start?: string | Date;
	time?: string;
	created?: string | Date;
	updated?: string | Date;
}

export interface EventseriesCollection {
	type: 'base';
	collectionId: string;
	collectionName: 'eventseries';
	response: EventseriesResponse;
	create: EventseriesCreate;
	update: EventseriesUpdate;
	relations: {
		events_via_series: EventsCollection[];
		team: TeamsCollection;
	};
}

// ==== end of eventseries block =====


// ===== Schema =====

export type Schema = {
	users: UsersCollection;
	events: EventsCollection;
	clubs: ClubsCollection;
	teams: TeamsCollection;
	requestcaches: RequestcachesCollection;
	leaguegroups: LeaguegroupsCollection;
	participations: ParticipationsCollection;
	uniformsets: UniformsetsCollection;
	_superusers: SuperusersCollection;
	_externalAuths: ExternalAuthsCollection;
	_mfas: MfasCollection;
	_otps: OtpsCollection;
	_authOrigins: AuthOriginsCollection;
	eventseries: EventseriesCollection;
}