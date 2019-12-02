package impl

import (
	. "strings"
	"unicode"

	def "github.com/emuggie/dipping/src/package/strings"
)

type delegator struct {
}

func (delegator) Compare(a, b string) int {
	return Compare(a, b)
}
func (delegator) Contains(s, substr string) bool {
	return Contains(s, substr)
}
func (delegator) ContainsAny(s, chars string) bool {
	return ContainsAny(s, chars)
}
func (delegator) ContainsRune(s string, r rune) bool {
	return ContainsRune(s, r)
}
func (delegator) Count(s, substr string) int {
	return Count(s, substr)
}
func (delegator) EqualFold(s, t string) bool {
	return EqualFold(s, t)
}
func (delegator) Fields(s string) []string {
	return Fields(s)
}
func (delegator) FieldsFunc(s string, f func(rune) bool) []string {
	return FieldsFunc(s, f)
}
func (delegator) HasPrefix(s, prefix string) bool {
	return HasPrefix(s, prefix)
}
func (delegator) HasSuffix(s, suffix string) bool {
	return HasSuffix(s, suffix)
}
func (delegator) Index(s, substr string) int {
	return Index(s, substr)
}
func (delegator) IndexAny(s, chars string) int {
	return IndexAny(s, chars)
}
func (delegator) IndexByte(s string, c byte) int {
	return IndexByte(s, c)
}
func (delegator) IndexFunc(s string, f func(rune) bool) int {
	return IndexFunc(s, f)
}
func (delegator) IndexRune(s string, r rune) int {
	return IndexRune(s, r)
}
func (delegator) Join(a []string, sep string) string {
	return Join(a, sep)
}
func (delegator) LastIndex(s, substr string) int {
	return LastIndex(s, substr)
}
func (delegator) LastIndexAny(s, chars string) int {
	return LastIndexAny(s, chars)
}
func (delegator) LastIndexByte(s string, c byte) int {
	return LastIndexByte(s, c)
}
func (delegator) LastIndexFunc(s string, f func(rune) bool) int {
	return LastIndexFunc(s, f)
}
func (delegator) Map(mapping func(rune) rune, s string) string {
	return Map(mapping, s)
}
func (delegator) Repeat(s string, count int) string {
	return Repeat(s, count)
}
func (delegator) Replace(s, old, new string, n int) string {
	return Replace(s, old, new, n)
}
func (delegator) ReplaceAll(s, old, new string) string {
	return ReplaceAll(s, old, new)
}
func (delegator) Split(s, sep string) []string {
	return Split(s, sep)
}
func (delegator) SplitAfter(s, sep string) []string {
	return SplitAfter(s, sep)
}
func (delegator) SplitAfterN(s, sep string, n int) []string {
	return SplitAfterN(s, sep, n)
}
func (delegator) SplitN(s, sep string, n int) []string {
	return SplitN(s, sep, n)
}
func (delegator) Title(s string) string {
	return Title(s)
}
func (delegator) ToLower(s string) string {
	return ToLower(s)
}
func (delegator) ToLowerSpecial(c unicode.SpecialCase, s string) string {
	return ToLowerSpecial(c, s)
}
func (delegator) ToTitle(s string) string {
	return ToTitle(s)
}
func (delegator) ToTitleSpecial(c unicode.SpecialCase, s string) string {
	return ToTitleSpecial(c, s)
}
func (delegator) ToUpper(s string) string {
	return ToUpper(s)
}
func (delegator) ToUpperSpecial(c unicode.SpecialCase, s string) string {
	return ToUpperSpecial(c, s)
}
func (delegator) ToValidUTF8(s, replacement string) string {
	return ToValidUTF8(s, replacement)
}
func (delegator) Trim(s string, cutset string) string {
	return Trim(s, cutset)
}
func (delegator) TrimFunc(s string, f func(rune) bool) string {
	return TrimFunc(s, f)
}
func (delegator) TrimLeft(s string, cutset string) string {
	return TrimLeft(s, cutset)
}
func (delegator) TrimLeftFunc(s string, f func(rune) bool) string {
	return TrimLeftFunc(s, f)
}
func (delegator) TrimPrefix(s, prefix string) string {
	return TrimPrefix(s, prefix)
}
func (delegator) TrimRight(s string, cutset string) string {
	return TrimRight(s, cutset)
}
func (delegator) TrimRightFunc(s string, f func(rune) bool) string {
	return TrimRightFunc(s, f)
}
func (delegator) TrimSpace(s string) string {
	return TrimSpace(s)
}
func (delegator) TrimSuffix(s, suffix string) string {
	return TrimSuffix(s, suffix)
}

var Delegator def.Strings

func init() {
	Delegator = new(delegator)
}

func Delegate() {
	def.Import = func() def.Strings {
		return Delegator
	}
}
