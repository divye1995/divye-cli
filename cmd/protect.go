/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/hex"
	"fmt"
	"os"

	"github.com/divye1995/divye-cli/util"
	"github.com/spf13/cobra"
)

// protectCmd represents the protect command
var protectCmd = &cobra.Command{
	Use:   "protect",
	Short: "encrypt and decrypt files",
	Long:  `encrypt and decrypt files using AES`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}
var outputFilePath string
var encryptCmd = &cobra.Command{
	Use:   "encrypt [path to file to protect] [passphrase]",
	Short: "encrypt files",
	Long: `generates encrypted file for given file using AES
		Example: encrypt supersecretfile.txt mykeyphrase 
	`,
	Args: cobra.ExactArgs(2),

	RunE: func(cmd *cobra.Command, args []string) error {
		filepath := args[0]
		key := args[1]
		// reading filepath
		dat, err := os.ReadFile(filepath)
		if err != nil {
			return err
		}
		// stringfile := string(dat)
		encryptedDat, err := util.Encrypt([]byte(key), dat)
		if err != nil {
			return err
		}
		if outputFilePath == "" {
			outputFilePath = fmt.Sprintf("%s.dsm", filepath)
		}
		// 0644 - The file's owner can read and write (6) Users in the same group as the file's owner can read (first 4) All users can read (second 4)

		err = os.WriteFile(outputFilePath, []byte(hex.EncodeToString(encryptedDat)), 0644)
		fmt.Println(filepath, " --> ", outputFilePath, " is protected")

		if err != nil {
			return err
		}
		return nil
	},
}

var decryptCmd = &cobra.Command{
	Use:   "decrypt [path to file to protect] [passphrase]",
	Short: "decrypt file",
	Long: ` decrypts a file that was encrypted using divye-cli
		Example: decrypt supersecretfile.txt.dsm passphrase 
	`,
	Args: cobra.ExactArgs(2),

	RunE: func(cmd *cobra.Command, args []string) error {
		filepath := args[0]
		key := args[1]
		// fmt.Println("File to be encrypted", filepath)
		// reading filepath
		dat, err := os.ReadFile(filepath)
		if err != nil {
			return err
		}

		toDecryptData, err := hex.DecodeString(string(dat))
		if err != nil {
			return nil
		}
		// stringfile := string(dat)
		decryptedDat, err := util.Decrypt([]byte(key), toDecryptData)
		if err != nil {
			return err
		}

		if outputFilePath == "" {
			outputFilePath = fmt.Sprintf("%s.dsm", filepath)
		}
		fmt.Println(string(decryptedDat))

		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	encryptCmd.Flags().StringVarP(&outputFilePath, "output", "o", "", "Output file to save the encrypted data")
	protectCmd.AddCommand(encryptCmd)
	protectCmd.AddCommand(decryptCmd)
	rootCmd.AddCommand(protectCmd)

}
