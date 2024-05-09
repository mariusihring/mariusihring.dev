import {REST, Routes} from "discord.js"
import {config} from "../config"
import {commands} from "./index";


const commandsData = Object.values(commands).map((command) => command.data)

const rest = new REST({version: "10"}).setToken(config.DISCORD_TOKEN as string);

type DeployCommandProps = {
    guildId: string,
}

export async function deployCommands({guildId}: DeployCommandProps) {
   try {
       console.log("Starting refreshing application (/) commands.");

       await rest.put(
           Routes.applicationGuildCommands(config.DISCORD_CLIENT_ID, guildId),
           {
               body: commandsData
           }
       );
       console.log("Succesfsfully reloaded application (/) commands.");
   }
   catch (error) {
       console.error(error);
   }
}