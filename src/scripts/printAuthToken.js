#!/usr/bin/env node

// @ts-nocheck
import PocketBase from 'pocketbase';
import dotenv from 'dotenv';

dotenv.config({
    path: "../../.env"
});

const pbUrl = process.env.PUBLIC_POCKETBASE_URL;
const adminEmail = process.env.POCKETBASE_EMAIL;
const adminPassword = process.env.POCKETBASE_PASSWORD;

const pb = new PocketBase(pbUrl);

async function authenticateAndPrint() {
    if (!adminEmail) {
        console.error("env not loaded")
        return
    }

    try {
        await pb.collection("_superusers").authWithPassword(adminEmail, adminPassword);
        console.log('Authenticated successfully!');
        console.log('Token:', pb.authStore.token);
    } catch (error) {
        console.error('Error:', error);
    }
}

await authenticateAndPrint();
