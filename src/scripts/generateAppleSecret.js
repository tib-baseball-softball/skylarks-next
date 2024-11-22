#!/usr/bin/env node

// @ts-nocheck
import PocketBase from 'pocketbase';
import dotenv from 'dotenv';

// wrote this before realising it can be done in the admin UI. Oh well. At least it works.

dotenv.config({
    path: "../../.env"
});

const pbUrl = process.env.PUBLIC_POCKETBASE_URL;
const adminEmail = process.env.ADMIN_EMAIL;
const adminPassword = process.env.ADMIN_PASSWORD;
const clientId = process.env.APPLE_CLIENT_ID;
const teamId = process.env.APPLE_TEAM_ID;
const keyId = process.env.APPLE_KEY_ID;
const privateKey = process.env.APPLE_PRIVATE_KEY;
const duration = parseInt(process.env.APPLE_KEY_DURATION, 10);

const pb = new PocketBase(pbUrl);

async function authenticateAndGenerateSecret() {
    if (!adminEmail) {
        console.error("env not loaded")
        return
    }

    try {
        await pb.collection("_superusers").authWithPassword(adminEmail, adminPassword);
        console.log('Authenticated successfully!');

        const secret = await pb.settings.generateAppleClientSecret(
            clientId,
            teamId,

            keyId,
            privateKey,
            duration
        );
        console.log('Apple Client Secret generated:', secret);
    } catch (error) {
        console.error('Error:', error);
    }
}

await authenticateAndGenerateSecret();
