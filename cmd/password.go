package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"math/rand"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate password",
	Long: `Generate password with customizable options.
For example:
passwordGenerator generate  -l 12 -d -s`,
	Run: generatePassword,
}

func init() {
	rootCmd.AddCommand(generateCmd)

	generateCmd.Flags().IntP("length", "l", 8, "Length of password")
	generateCmd.Flags().BoolP("digits", "d", false, "Include digits in password")
	generateCmd.Flags().BoolP("special-char", "s", false, "Include special characters in password")
}

func generatePassword(cmd *cobra.Command, args []string) {
	fmt.Println("Generating Password...")

	length, _ := cmd.Flags().GetInt("length")
	isDigit, _ := cmd.Flags().GetBool("digits")
	isSpecialChar, _ := cmd.Flags().GetBool("special-char")

	charset := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

	if isDigit {
		charset += "0123456789"
	}
	if isSpecialChar {
		charset += "!@#$%^&*()_+{}[]|;:<>,.?-="
	}

	password := make([]byte, length)

	for i := range password {
		password[i] = charset[rand.Intn(len(charset))]
	}

	fmt.Println(string(password))
}
