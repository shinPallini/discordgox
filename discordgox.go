package discordgox

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

type CommandSettings struct {
	Commands          []*discordgo.ApplicationCommand
	CommandHandlers   map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate)
	ComponentHandlers map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate)
}

func (s *CommandSettings) AddCommand(command *discordgo.ApplicationCommand, fn func(s *discordgo.Session, i *discordgo.InteractionCreate)) {
	_, exist := DefaultSettings.CommandHandlers[command.Name]
	if exist {
		log.Fatal(fmt.Sprintf("[%s] ← このコマンド名が重複しています！", command.Name))
	}
	// コマンド部分のNameをそのままmapのKeyとして設定しておく
	s.CommandHandlers[command.Name] = fn
	s.Commands = append(s.Commands, command)
}

func (s *CommandSettings) AddComponent(customID string, fn func(s *discordgo.Session, i *discordgo.InteractionCreate)) {
	_, exist := DefaultSettings.ComponentHandlers[customID]
	if exist {
		log.Fatal(fmt.Sprintf("[%s] ← このカスタムIDが重複しています！", customID))
	}
	s.ComponentHandlers[customID] = fn
}

func (s *CommandSettings) AddCommandWithComponent(
	cmd *discordgo.ApplicationCommand,
	cmdfn func(s *discordgo.Session, i *discordgo.InteractionCreate),
	customID string,
	cpnfn func(s *discordgo.Session, i *discordgo.InteractionCreate),
) {
	s.AddComponent(customID, cpnfn)
	s.AddCommand(cmd, cmdfn)
}

var (
	DefaultSettings = CommandSettings{
		Commands:          make([]*discordgo.ApplicationCommand, 0),
		CommandHandlers:   make(map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate)),
		ComponentHandlers: make(map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate)),
	}
)

func AddCommand(command *discordgo.ApplicationCommand, fn func(s *discordgo.Session, i *discordgo.InteractionCreate)) {
	_, exist := DefaultSettings.CommandHandlers[command.Name]
	if exist {
		log.Fatal(fmt.Sprintf("[%s] ← このコマンド名が重複しています！", command.Name))
	}
	// コマンド部分のNameをそのままmapのKeyとして設定しておく
	DefaultSettings.CommandHandlers[command.Name] = fn
	DefaultSettings.Commands = append(DefaultSettings.Commands, command)
}

func AddComponent(customID string, fn func(s *discordgo.Session, i *discordgo.InteractionCreate)) {
	_, exist := DefaultSettings.ComponentHandlers[customID]
	if exist {
		log.Fatal(fmt.Sprintf("[%s] ← このカスタムIDが重複しています！", customID))
	}
	DefaultSettings.ComponentHandlers[customID] = fn
}

func AddCommandWithComponent(
	cmd *discordgo.ApplicationCommand,
	cmdfn func(s *discordgo.Session, i *discordgo.InteractionCreate),
	customID string,
	cpnfn func(s *discordgo.Session, i *discordgo.InteractionCreate),
) {
	DefaultSettings.AddComponent(customID, cpnfn)
	DefaultSettings.AddCommand(cmd, cmdfn)
}

// 以下Option型と構造体生成関数の記述
// InteractionResponse構造体の初期化のためのOptionと関数
type InteractionResponseOption func(r *discordgo.InteractionResponse)

func SetType(t discordgo.InteractionResponseType) InteractionResponseOption {
	return func(r *discordgo.InteractionResponse) {
		r.Type = t
	}
}

func SetData(rd *discordgo.InteractionResponseData) InteractionResponseOption {
	return func(r *discordgo.InteractionResponse) {
		r.Data = rd
	}
}

func NewInteractionResponse(options ...InteractionResponseOption) *discordgo.InteractionResponse {
	ir := &discordgo.InteractionResponse{}

	for _, opt := range options {
		opt(ir)
	}
	return ir
}

// InteractionResponseData構造体の初期化のためのOptionと関数
type InteractionRsponseDataOption func(rd *discordgo.InteractionResponseData)

func SetContent(content string) InteractionRsponseDataOption {
	return func(rd *discordgo.InteractionResponseData) {
		rd.Content = content
	}
}

func SetEmbed(e []*discordgo.MessageEmbed) InteractionRsponseDataOption {
	return func(rd *discordgo.InteractionResponseData) {
		rd.Embeds = append(rd.Embeds, e...)
	}
}

func SetComponent(c []discordgo.MessageComponent) InteractionRsponseDataOption {
	return func(rd *discordgo.InteractionResponseData) {
		rd.Components = append(rd.Components, c...)
	}
}

func NewInteractionResponseData(options ...InteractionRsponseDataOption) *discordgo.InteractionResponseData {
	ird := &discordgo.InteractionResponseData{}

	for _, opt := range options {
		opt(ird)
	}
	return ird
}

// MessageEmbed構造体の初期化のためのOptionと関数
type MessageEmbedOption func(e *discordgo.MessageEmbed)

func SetEmbedType(t discordgo.EmbedType) MessageEmbedOption {
	return func(e *discordgo.MessageEmbed) {
		e.Type = t
	}
}

func SetTitle(s string) MessageEmbedOption {
	return func(e *discordgo.MessageEmbed) {
		e.Title = s
	}
}

func SetDescription(s string) MessageEmbedOption {
	return func(e *discordgo.MessageEmbed) {
		e.Description = s
	}
}

// colorは16進数で指定する
func SetColor(i int) MessageEmbedOption {
	return func(e *discordgo.MessageEmbed) {
		e.Color = i
	}
}

func SetEmbedField(ef []*discordgo.MessageEmbedField) MessageEmbedOption {
	return func(e *discordgo.MessageEmbed) {
		e.Fields = append(e.Fields, ef...)
	}
}

func NewMessageEmbed(options ...MessageEmbedOption) *discordgo.MessageEmbed {
	e := &discordgo.MessageEmbed{}

	for _, opt := range options {
		opt(e)
	}
	return e
}

// MessageEmbedField構造体の初期化のためのOptionと関数
type MessageEmbedFieldOption func(ef *discordgo.MessageEmbedField)

func SetEmbedFieldName(s string) MessageEmbedFieldOption {
	return func(ef *discordgo.MessageEmbedField) {
		ef.Name = s
	}
}

func SetEmbedFieldValue(s string) MessageEmbedFieldOption {
	return func(ef *discordgo.MessageEmbedField) {
		ef.Value = s
	}
}

func SetEmbedFieldInline(b bool) MessageEmbedFieldOption {
	return func(ef *discordgo.MessageEmbedField) {
		ef.Inline = b
	}
}

func NewMessageEmbedField(options ...MessageEmbedFieldOption) *discordgo.MessageEmbedField {
	ef := &discordgo.MessageEmbedField{}

	for _, opt := range options {
		opt(ef)
	}
	return ef
}

// MessageComponent構造体初期化のためのOptionと関数
type ActionsRowOption func(*discordgo.ActionsRow)

func SetLinkButton(label string, url string) ActionsRowOption {
	return func(r *discordgo.ActionsRow) {
		r.Components = append(r.Components, discordgo.Button{
			Style: discordgo.LinkButton,
			Label: label,
			URL:   url,
		})
	}
}

func SetCustomButton(style discordgo.ButtonStyle, label string, customID string) ActionsRowOption {
	return func(r *discordgo.ActionsRow) {
		r.Components = append(r.Components, discordgo.Button{
			Style:    style,
			Label:    label,
			CustomID: customID,
		})
	}
}

func SetSingleSelectMenu(customID string, selectOptions []discordgo.SelectMenuOption) ActionsRowOption {
	return func(r *discordgo.ActionsRow) {
		r.Components = append(r.Components, discordgo.SelectMenu{
			CustomID: customID,
			Options:  selectOptions,
		})
	}
}

func SetMultiSelectMenu(customID string, selectOptions []discordgo.SelectMenuOption, min_value *int, max_value int) ActionsRowOption {
	return func(r *discordgo.ActionsRow) {
		r.Components = append(r.Components, discordgo.SelectMenu{
			CustomID:  customID,
			Options:   selectOptions,
			MinValues: min_value,
			MaxValues: max_value,
		})
	}
}

func NewActionsRow(options ...ActionsRowOption) *discordgo.ActionsRow {
	c := &discordgo.ActionsRow{}

	for _, opt := range options {
		opt(c)
	}
	return c
}

// SelectMenuOption構造体初期化のためのOptionと関数
type SelectMenuOptionOption func(*discordgo.SelectMenuOption)

func SetSelectDescription(d string) SelectMenuOptionOption {
	return func(o *discordgo.SelectMenuOption) {
		o.Description = d
	}
}

func SetSelectDefaultEmoji(s string) SelectMenuOptionOption {
	return func(o *discordgo.SelectMenuOption) {
		o.Emoji = discordgo.ComponentEmoji{
			Name: s,
		}
	}
}

func SetSelectCustomEmoji(s string, id string) SelectMenuOptionOption {
	return func(o *discordgo.SelectMenuOption) {
		o.Emoji = discordgo.ComponentEmoji{
			Name: s,
			ID:   id,
		}
	}
}

func NewSelectMenuOption(label string, value string, options ...SelectMenuOptionOption) *discordgo.SelectMenuOption {
	o := &discordgo.SelectMenuOption{
		Label: label,
		Value: value,
	}

	for _, opt := range options {
		opt(o)
	}
	return o
}

// あらゆる型の変数を引数に取って配列を返すメソッド
// interfaceの型が必要な場合はインスタンス化するためのかっこの中に型を指定する
func NewList[T any](t ...T) []T {
	l := []T{}
	for _, elem := range t {
		l = append(l, elem)
	}
	return l
}
