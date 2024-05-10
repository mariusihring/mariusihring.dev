import dotenv from "dotenv";

dotenv.config();

const {DISCORD_TOKEN, DISCORD_CLIENT_ID, WAIFU_TOKEN} = process.env;

if (!DISCORD_CLIENT_ID || !DISCORD_CLIENT_ID || !WAIFU_TOKEN) {
    throw new Error("environment variable is missing");
}


export const config = {
    DISCORD_TOKEN,
    DISCORD_CLIENT_ID,
    WAIFU_TOKEN
}