// This file was auto-generated from the
// Grooveshark API Extractor
package main


import (
	// TODO add imports
)

type GoSharky struct {
	// TODO impl
}

//  Use addUserLibrarySongsEx instead. Add songs to a user's library. Song metadata should be spread across all 3 params. albumIDs[0] should be the respective albumID for songIDs[0] and same with artistIDs.
// Note: You must provide a sessionID with this method.
func (sharky *GoSharky) AddUserLibrarySongs(songIDs, albumIDs, artistIDs string) {
	// TODO impelemnt
}

// Get user library songs. Requires an authenticated session.
// Note: You must provide a sessionID with this method.
func (sharky *GoSharky) GetUserLibrarySongs(limit, page int) {
	// TODO impelemnt
}

// Add songs to a user's library. Songs should be an array of objects representing each song with keys: songID, albumID, artistID, trackNum.
// Note: You must provide a sessionID with this method.
func (sharky *GoSharky) AddUserLibrarySongsEx(songs string) {
	// TODO impelemnt
}

// Remove songs from a user's library.
// Note: You must provide a sessionID with this method.
func (sharky *GoSharky) RemoveUserLibrarySongs(songIDs, albumIDs, artistIDs string) {
	// TODO impelemnt
}

// Get subscribed playlists of the logged-in user. Requires an authenticated session.
// Note: You must provide a sessionID with this method.
func (sharky *GoSharky) GetUserPlaylistsSubscribed() {
	// TODO impelemnt
}

// Get playlists of the logged-in user. Requires an authenticated session.
// Note: You must provide a sessionID with this method.
func (sharky *GoSharky) GetUserPlaylists(limit int) {
	// TODO impelemnt
}

// Get user favorite songs. Requires an authenticated session.
// Note: You must provide a sessionID with this method.
func (sharky *GoSharky) GetUserFavoriteSongs(limit int) {
	// TODO impelemnt
}

// Remove a set of favorite songs for a user. Must provide a logged-in sessionID.
// Note: You must provide a sessionID with this method.
func (sharky *GoSharky) RemoveUserFavoriteSongs(songIDs string) {
	// TODO impelemnt
}

// Logout a user using an established session.
// Note: You must provide a sessionID with this method.
func (sharky *GoSharky) Logout() {
	// TODO impelemnt
}

// Authenticate a user using a token from http://grooveshark.com/auth/. See Overview for documentation.
// Note: You must provide a sessionID with this method.
func (sharky *GoSharky) AuthenticateToken(token string) {
	// TODO impelemnt
}

// Get logged-in user info from sessionID
// Note: You must provide a sessionID with this method.
func (sharky *GoSharky) GetUserInfo() {
	// TODO impelemnt
}

// Get logged-in user subscription info. Returns type of subscription and either dateEnd or recurring.
// Note: You must provide a sessionID with this method.
func (sharky *GoSharky) GetUserSubscriptionDetails() {
	// TODO impelemnt
}

// Add a favorite song for a user. Must provide a logged-in sessionID.
// Note: You must provide a sessionID with this method.
func (sharky *GoSharky) AddUserFavoriteSong(songID int) {
	// TODO impelemnt
}

// Subscribe to a playlist for the logged-in user. Requires an authenticated session.
// Note: You must provide a sessionID with this method.
func (sharky *GoSharky) SubscribePlaylist(playlistID int) {
	// TODO impelemnt
}

// Unsubscribe from a playlist for the logged-in user. Requires an authenticated session.
// Note: You must provide a sessionID with this method.
func (sharky *GoSharky) UnsubscribePlaylist(playlistID int) {
	// TODO impelemnt
}

// Get country from IP. If an IP is omitted, it will use the request's IP.
func (sharky *GoSharky) GetCountry(ip string) {
	// TODO impelemnt
}

// Get playlist information. To get songs as well, call getPlaylist.
func (sharky *GoSharky) GetPlaylistInfo(playlistID string) {
	// TODO impelemnt
}

// Get a subset of today's popular songs, from the Grooveshark popular billboard.
func (sharky *GoSharky) GetPopularSongsToday(limit int) {
	// TODO impelemnt
}

// Get a subset of this month's popular songs, from the Grooveshark popular billboard.
func (sharky *GoSharky) GetPopularSongsMonth(limit int) {
	// TODO impelemnt
}

// Useful for testing if the service is up. Returns "Hello, World" in various languages.
func (sharky *GoSharky) PingService() {
	// TODO impelemnt
}

// Describe service methods
func (sharky *GoSharky) GetServiceDescription() {
	// TODO impelemnt
}

// Undeletes a playlist.
// Note: You must provide a sessionID with this method.
func (sharky *GoSharky) UndeletePlaylist(playlistID int) {
	// TODO impelemnt
}

// Deletes a playlist.
// Note: You must provide a sessionID with this method.
func (sharky *GoSharky) DeletePlaylist(playlistID int) {
	// TODO impelemnt
}

// Get songs on a playlist. Use getPlaylist instead.
func (sharky *GoSharky) GetPlaylistSongs(playlistID string, limit int) {
	// TODO impelemnt
}

// Get playlist info and songs.
func (sharky *GoSharky) GetPlaylist(playlistID string, limit int) {
	// TODO impelemnt
}

// Set playlist songs, overwrites any already saved
// Note: You must provide a sessionID with this method.
func (sharky *GoSharky) SetPlaylistSongs(playlistID int, songIDs string) {
	// TODO impelemnt
}

// Create a new playlist, optionally adding songs to it.
// Note: You must provide a sessionID with this method.
func (sharky *GoSharky) CreatePlaylist(name, songIDs string) {
	// TODO impelemnt
}

// Renames a playlist.
// Note: You must provide a sessionID with this method.
func (sharky *GoSharky) RenamePlaylist(playlistID int, name string) {
	// TODO impelemnt
}

// Authenticate a user using an established session, a login and an md5 of their password.
// Note: You must provide a sessionID with this method.
func (sharky *GoSharky) Authenticate(login, password string) {
	// TODO impelemnt
}

// Get userID from username
func (sharky *GoSharky) GetUserIDFromUsername(username string) {
	// TODO impelemnt
}

// Get meta-data information about one or more albums
func (sharky *GoSharky) GetAlbumsInfo(albumIDs string) {
	// TODO impelemnt
}

// Get songs on an album. Returns all songs, verified and unverified
func (sharky *GoSharky) GetAlbumSongs(albumID, limit int) {
	// TODO impelemnt
}

// Get meta-data information about one or more artists
func (sharky *GoSharky) GetArtistsInfo(artistIDs string) {
	// TODO impelemnt
}

// Get information about a song or multiple songs.  The songID(s) should always be passed in as an array.
func (sharky *GoSharky) GetSongsInfo(songIDs string) {
	// TODO impelemnt
}

// Check if an album exists
func (sharky *GoSharky) GetDoesAlbumExist(albumID int) {
	// TODO impelemnt
}

// Check if a song exists
func (sharky *GoSharky) GetDoesSongExist(songID int) {
	// TODO impelemnt
}

// Check if an artist exists
func (sharky *GoSharky) GetDoesArtistExist(artistID int) {
	// TODO impelemnt
}

// Authenticate a user (login) using an established session. Please use the authenticate method instead.
// Note: You must provide a sessionID with this method.
func (sharky *GoSharky) AuthenticateUser(username, token string) {
	// TODO impelemnt
}

// Get an artist's verified albums
func (sharky *GoSharky) GetArtistVerifiedAlbums(artistID int) {
	// TODO impelemnt
}

// Get an artist's albums, verified and unverified
func (sharky *GoSharky) GetArtistAlbums(artistID int) {
	// TODO impelemnt
}

// Get 100 popular songs for an artist
func (sharky *GoSharky) GetArtistPopularSongs(artistID int) {
	// TODO impelemnt
}

// ================= Search =================

// Perform a playlist search.
func (sharky *GoSharky) GetPlaylistSearchResults(query string, limit int) {
	// TODO impelemnt
}

// Perform an album search.
func (sharky *GoSharky) GetAlbumSearchResults(query string, limit int) {
	// TODO impelemnt
}

// Perform a song search.
func (sharky *GoSharky) GetSongSearchResults(query string, country object, limit, offset int) {
	// TODO impelemnt
}

// Perform an artist search.
func (sharky *GoSharky) GetArtistSearchResults(query string, limit int) {
	// TODO impelemnt
}

// ================= Streams =================

// Get stream key, ID, etc. from songID. Requires country object obtained from getCountry
// Note: You must provide a sessionID with this method.
func (sharky *GoSharky) GetStreamKeyStreamServer(songID int, country object, lowBitrate bool) {
	// TODO impelemnt
}

// ================= URLS =================

// Get Grooveshark URL for tinysong base 62.
func (sharky *GoSharky) GetSongURLFromTinysongBase62(base62 string) {
	// TODO impelemnt
}

// Get playable song URL from songID
func (sharky *GoSharky) GetSongURLFromSongID(songID int) {
	// TODO impelemnt
}

// Get playlist URL from playlistID
func (sharky *GoSharky) GetPlaylistURLFromPlaylistID(playlistID int) {
	// TODO impelemnt
}

// Get a song's Tinysong.com url.
func (sharky *GoSharky) GetTinysongURLFromSongID(songID int) {
	// TODO impelemnt
}

// ================= Users (no auth) =================

// Get playlists created by a userID. Does not require an authenticated session.
func (sharky *GoSharky) GetUserPlaylistsByUserID(userID, limit int) {
	// TODO impelemnt
}

// Get user info from userID
func (sharky *GoSharky) GetUserInfoFromUserID(userID int) {
	// TODO impelemnt
}

// ================= Recs =================

// Get similar artist for a given artistID.
func (sharky *GoSharky) GetSimilarArtists(artistID, limit, page int) {
	// TODO impelemnt
}

// ================= Sessions =================

// Start a session
func (sharky *GoSharky) StartSession() {
	// TODO impelemnt
}

// ================= Trials =================

// Gets a trial for an application and the provided uniqueID or logged in user.
// Note: You must provide a sessionID with this method.
func (sharky *GoSharky) GetTrialInfo(uniqueID string) {
	// TODO impelemnt
}

// Starts a trial for a user bound to your application and the provided uniqueID.
// Note: You must provide a sessionID with this method.
func (sharky *GoSharky) CreateTrial(uniqueID string) {
	// TODO impelemnt
}

// ================= Autocomplete =================

// Autocomplete search. Type parameter is 'music', 'playlist', or 'user'. Returns an array of words.
func (sharky *GoSharky) GetAutocompleteSearchResults(query, type string, limit int) {
	// TODO impelemnt
}

// ================= Subscriber streams =================

// Get stream key, ID, etc. from songID for a subscriber account.  Requires country object obtained from getCountry and a logged-in sessionID from a Grooveshark Anywhere subscriber.
// Note: You must provide a sessionID with this method.
func (sharky *GoSharky) GetSubscriberStreamKey(songID int, country object, lowBitrate bool, uniqueID string) {
	// TODO impelemnt
}

// Mark a song as having been played for greater than or equal to 30 seconds.
// Note: You must provide a sessionID with this method.
func (sharky *GoSharky) MarkStreamKeyOver30Secs(streamKey string, streamServerID int, uniqueID string) {
	// TODO impelemnt
}

// ================= Subscriber streams =================

// Mark a song as complete (played for greater than or equal to 30 seconds, and having reached the last second either through seeking or normal playback).
// Note: You must provide a sessionID with this method.
func (sharky *GoSharky) MarkSongComplete(songID int, streamKey string, streamServerID int, autoplayState object) {
	// TODO impelemnt
}

// ================= Autoplay =================

// Grab a relevant song for autoplay
func (sharky *GoSharky) GetAutoplaySong(autoplayState object) {
	// TODO impelemnt
}

// Gets a list of tags (stations)
func (sharky *GoSharky) GetAutoplayTags() {
	// TODO impelemnt
}

// Start autoplay using a tag and grab a relevant song
func (sharky *GoSharky) StartAutoplayTag(tagID int) {
	// TODO impelemnt
}

// Start autoplay and grab a relevant song
func (sharky *GoSharky) StartAutoplay(artistIDs, songIDs object) {
	// TODO impelemnt
}

// Remove a vote up for a song
func (sharky *GoSharky) RemoveVoteUpAutoplaySong(song, autoplayState object) {
	// TODO impelemnt
}

// Vote up a song
func (sharky *GoSharky) VoteUpAutoplaySong(song, autoplayState object) {
	// TODO impelemnt
}

// Remove a song from the autoplay state
func (sharky *GoSharky) RemoveSongFromAutoplay(song, autoplayState object) {
	// TODO impelemnt
}

// Add a song to the autoplay state
func (sharky *GoSharky) AddSongToAutoplay(song, autoplayState object) {
	// TODO impelemnt
}

// Vote down a song
func (sharky *GoSharky) VoteDownAutoplaySong(song, autoplayState object) {
	// TODO impelemnt
}

// Remove a vote down for a song
func (sharky *GoSharky) RemoveVoteDownAutoplaySong(song, autoplayState object) {
	// TODO impelemnt
}

// ================= Tinysong =================

// Get Grooveshark songID for tinysong base 62.
func (sharky *GoSharky) GetSongIDFromTinysongBase62(base62 string) {
	// TODO impelemnt
}

// ================= Register =================

// Register and authenticate a user using an established session. The username is alpha-numeric with a period, dash or underscore allowed in the middle. The username can be blank or 5-32 characters.  Passwords must be between 5 and 32 characters.
// Note: You must provide a sessionID with this method.
func (sharky *GoSharky) RegisterUser(emailAddress, password, fullName, username, gender, birthDate string) {
	// TODO impelemnt
}
