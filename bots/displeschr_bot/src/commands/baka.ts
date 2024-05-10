import {CommandInteraction, SlashCommandBuilder} from "discord.js";

export const data = new SlashCommandBuilder()
    .setName("baka")
    .setDescription("Replies with a random baka GIF!");

export async function execute(interaction: CommandInteraction) {
          // return interaction.reply({files: [data.data.url]})
    if (interaction.user.dmChannel) {
        await interaction.user.dmChannel.send("test hehehhe")
    } else {
        const dmChannel = await interaction.user.createDM();
       await dmChannel.send("test hehehhe")
    }
    return interaction.reply("done")
}