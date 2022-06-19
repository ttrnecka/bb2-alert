package bb2

import "encoding/xml"

const MESSAGE_NAME_MATCH_FOUND = "NotificationAutomatchOpponentFound"

const HEADER_START = "<Header>"
const HEADER_END = "</Header>"

type XmlHeader struct {
	XMLName xml.Name `xml:"Header"`
	Data    Data     `xml:"Data"`
}

type Data struct {
	XMLName               xml.Name `xml:"Data"`
	Type                  string   `xml:"type,attr"`
	Size                  int      `xml:"size,attr"`
	Zipped                string   `xml:"zipped,attr"`
	MessageName           string   `xml:"MessageName,attr"`
	MessageToken          string   `xml:"MessageToken,attr"`
	SizeBeforeCompression int      `xml:"sizeBeforeCompression,attr"`
}

type NotificationAutomatchOpponentFound struct {
	XMLName      xml.Name `xml:"NotificationAutomatchOpponentFound"`
	Text         string   `xml:",chardata"`
	Map          string   `xml:"Map"`
	OwnedDLCAway string   `xml:"OwnedDLCAway"`
	OpponentId   string   `xml:"OpponentId"`
	SessionId    string   `xml:"SessionId"`
	OpponentName string   `xml:"OpponentName"`
	Specific     struct {
		Text        string `xml:",chardata"`
		Name        string `xml:"Name"`
		MessageData struct {
			Text          string `xml:",chardata"`
			BBGameSession struct {
				Text         string `xml:",chardata"`
				OpponentName string `xml:"OpponentName"`
				OpponentAi   string `xml:"OpponentAi"`
				TeamRosters  struct {
					Text       string `xml:",chardata"`
					TeamRoster []struct {
						Text             string `xml:",chardata"`
						CoachProgression struct {
							Text                string `xml:",chardata"`
							RowCoachProgression struct {
								Text                       string `xml:",chardata"`
								IdCoach                    string `xml:"IdCoach"`
								LevelsSinceLastTalentPoint string `xml:"LevelsSinceLastTalentPoint"`
								Level                      string `xml:"Level"`
								SponsorshipPoints          string `xml:"SponsorshipPoints"`
								TalentPointsSpent          string `xml:"TalentPointsSpent"`
								TalentPoints               string `xml:"TalentPoints"`
								Portrait                   string `xml:"Portrait"`
								Xp                         string `xml:"Xp"`
								ID                         string `xml:"Id"`
							} `xml:"RowCoachProgression"`
							XpForNextLevel    string `xml:"XpForNextLevel"`
							XpForCurrentLevel string `xml:"XpForCurrentLevel"`
						} `xml:"CoachProgression"`
						PlayerInfos struct {
							Text        string `xml:",chardata"`
							PlayerInfos []struct {
								Text   string `xml:",chardata"`
								Skills string `xml:"Skills"`
								Player struct {
									Text string `xml:",chardata"`
									Row  struct {
										Text                     string `xml:",chardata"`
										IdPlayerLevels           string `xml:"IdPlayerLevels"`
										IdRaces                  int    `xml:"IdRaces"`
										Number                   string `xml:"Number"`
										Dead                     string `xml:"Dead"`
										Experience               string `xml:"Experience"`
										Retired                  string `xml:"Retired"`
										LevelUpDouble            string `xml:"LevelUpDouble"`
										CharacsMovementAllowance string `xml:"CharacsMovementAllowance"`
										MatchSuspended           string `xml:"MatchSuspended"`
										NbMatchsSinceAgeRoll     string `xml:"NbMatchsSinceAgeRoll"`
										Biography                string `xml:"Biography"`
										NbLevelsUp               string `xml:"NbLevelsUp"`
										LevelUpRollResult        string `xml:"LevelUpRollResult"`
										Star                     string `xml:"Star"`
										CharacsArmourValue       string `xml:"CharacsArmourValue"`
										IdTeamListing            struct {
											Text  string `xml:",chardata"`
											Value string `xml:"Value"`
										} `xml:"IdTeamListing"`
										CharacsAgility string `xml:"CharacsAgility"`
										ID             struct {
											Text  string `xml:",chardata"`
											Value string `xml:"Value"`
										} `xml:"ID"`
										Name               string `xml:"Name"`
										CharacsStrength    string `xml:"CharacsStrength"`
										Age                string `xml:"Age"`
										LevelUpRollResult2 string `xml:"LevelUpRollResult2"`
										Value              string `xml:"Value"`
										IdHead             string `xml:"IdHead"`
										IdPlayerTypes      string `xml:"IdPlayerTypes"`
									} `xml:"Row"`
								} `xml:"Player"`
								Statistics struct {
									Text             string `xml:",chardata"`
									StatisticsPlayer []struct {
										Text     string `xml:",chardata"`
										Category string `xml:"Category"`
										Row      struct {
											Text                   string `xml:",chardata"`
											InflictedCasualties    string `xml:"InflictedCasualties"`
											InflictedStuns         string `xml:"InflictedStuns"`
											InflictedPasses        string `xml:"InflictedPasses"`
											SustainedInterceptions string `xml:"SustainedInterceptions"`
											InflictedMetersPassing string `xml:"InflictedMetersPassing"`
											Category               string `xml:"Category"`
											InflictedTackles       string `xml:"InflictedTackles"`
											SustainedTackles       string `xml:"SustainedTackles"`
											InflictedKO            string `xml:"InflictedKO"`
											SustainedInjuries      string `xml:"SustainedInjuries"`
											SustainedKO            string `xml:"SustainedKO"`
											SustainedDead          string `xml:"SustainedDead"`
											MatchPlayed            string `xml:"MatchPlayed"`
											InflictedDead          string `xml:"InflictedDead"`
											MVP                    string `xml:"MVP"`
											InflictedInterceptions string `xml:"InflictedInterceptions"`
											InflictedPushOuts      string `xml:"InflictedPushOuts"`
											InflictedCatches       string `xml:"InflictedCatches"`
											SustainedCasualties    string `xml:"SustainedCasualties"`
											InflictedInjuries      string `xml:"InflictedInjuries"`
											ID                     string `xml:"ID"`
											DicesStat              struct {
												Text         string `xml:",chardata"`
												BlockDistrib string `xml:"BlockDistrib"`
												Roll         string `xml:"Roll"`
												Block        string `xml:"Block"`
												RollDistrib  string `xml:"RollDistrib"`
											} `xml:"DicesStat"`
											SustainedStuns         string `xml:"SustainedStuns"`
											InflictedMetersRunning string `xml:"InflictedMetersRunning"`
											InflictedTouchdowns    string `xml:"InflictedTouchdowns"`
											IdPlayerListing        string `xml:"IdPlayerListing"`
										} `xml:"Row"`
									} `xml:"StatisticsPlayer"`
								} `xml:"Statistics"`
								Casualties   string `xml:"Casualties"`
								SellingPrice string `xml:"SellingPrice"`
							} `xml:"PlayerInfos"`
						} `xml:"PlayerInfos"`
						Competitions struct {
							Text                string `xml:",chardata"`
							TeamCompetitionData struct {
								Text        string `xml:",chardata"`
								LeagueOwner struct {
									Text      string `xml:",chardata"`
									IdUser    string `xml:"IdUser"`
									Status    string `xml:"Status"`
									Groupmask string `xml:"Groupmask"`
									User      string `xml:"User"`
									Lastlogin string `xml:"Lastlogin"`
								} `xml:"LeagueOwner"`
								CompetitionOwner struct {
									Text      string `xml:",chardata"`
									IdUser    string `xml:"IdUser"`
									Status    string `xml:"Status"`
									Groupmask string `xml:"Groupmask"`
									User      string `xml:"User"`
									Lastlogin string `xml:"Lastlogin"`
								} `xml:"CompetitionOwner"`
								RowLeague struct {
									Text          string `xml:",chardata"`
									EmailContacts string `xml:"EmailContacts"`
									Edition       string `xml:"Edition"`
									IdOwner       string `xml:"IdOwner"`
									LeagueType    string `xml:"LeagueType"`
									Logo          string `xml:"Logo"`
									ID            struct {
										Text  string `xml:",chardata"`
										Value string `xml:"Value"`
									} `xml:"Id"`
									MarketPlaceRookieGenerated string `xml:"MarketPlaceRookieGenerated"`
									HasPassword                string `xml:"HasPassword"`
									Solo                       string `xml:"Solo"`
									Description                string `xml:"Description"`
									MarketPlace                string `xml:"MarketPlace"`
									NeedValidation             string `xml:"NeedValidation"`
									Phase                      string `xml:"Phase"`
									Password                   string `xml:"Password"`
									Data                       string `xml:"Data"`
									Name                       string `xml:"Name"`
									Official                   string `xml:"Official"`
									Websites                   string `xml:"Websites"`
									TreasurySp                 string `xml:"TreasurySp"`
									NbRegisteredTeams          string `xml:"NbRegisteredTeams"`
									Flags                      string `xml:"Flags"`
								} `xml:"RowLeague"`
								LocalLeague string `xml:"LocalLeague"`
								IdLocalTeam struct {
									Text  string `xml:",chardata"`
									Value string `xml:"Value"`
								} `xml:"IdLocalTeam"`
								IdTeam struct {
									Text  string `xml:",chardata"`
									Value string `xml:"Value"`
								} `xml:"IdTeam"`
								RowCompetition struct {
									Text              string `xml:",chardata"`
									IdPreviousEdition struct {
										Text  string `xml:",chardata"`
										Value string `xml:"Value"`
									} `xml:"IdPreviousEdition"`
									IdOwner                string `xml:"IdOwner"`
									RankedChallengeAllowed string `xml:"RankedChallengeAllowed"`
									Logo                   string `xml:"Logo"`
									RankedMatchAllowed     string `xml:"RankedMatchAllowed"`
									CheckinPrice           string `xml:"CheckinPrice"`
									ID                     struct {
										Text  string `xml:",chardata"`
										Value string `xml:"Value"`
									} `xml:"Id"`
									AutomatchAllowed         string `xml:"AutomatchAllowed"`
									KeepReplayPolicy         string `xml:"KeepReplayPolicy"`
									CloneOnNextLeagueEdition string `xml:"CloneOnNextLeagueEdition"`
									LeaguePhase              string `xml:"LeaguePhase"`
									ChallengeAllowed         string `xml:"ChallengeAllowed"`
									IdLeague                 struct {
										Text  string `xml:",chardata"`
										Value string `xml:"Value"`
									} `xml:"IdLeague"`
									NbRegisteredTeams        string `xml:"NbRegisteredTeams"`
									AllowSelfMatchValidation string `xml:"AllowSelfMatchValidation"`
									PrematchEvents           string `xml:"PrematchEvents"`
									Flags                    string `xml:"Flags"`
									Rewards                  string `xml:"Rewards"`
									Description              string `xml:"Description"`
									IdCompetitionTypes       string `xml:"IdCompetitionTypes"`
									RegistrationMode         string `xml:"RegistrationMode"`
									NbRounds                 string `xml:"NbRounds"`
									NbTeamsMax               string `xml:"NbTeamsMax"`
									CompetitionStatus        string `xml:"CompetitionStatus"`
									AdmissionStadiumLevel    string `xml:"AdmissionStadiumLevel"`
									UnRankedChallengeAllowed string `xml:"UnRankedChallengeAllowed"`
									Name                     string `xml:"Name"`
									CurrentRound             string `xml:"CurrentRound"`
									LeagueEdition            string `xml:"LeagueEdition"`
									UnRankedMatchAllowed     string `xml:"UnRankedMatchAllowed"`
									PureAutomatch            string `xml:"PureAutomatch"`
									NeedValidation           string `xml:"NeedValidation"`
									AcceptTicketRequest      string `xml:"AcceptTicketRequest"`
									TurnDuration             string `xml:"TurnDuration"`
									GameOptions              string `xml:"GameOptions"`
								} `xml:"RowCompetition"`
							} `xml:"TeamCompetitionData"`
						} `xml:"Competitions"`
						RowLeague struct {
							Text          string `xml:",chardata"`
							EmailContacts string `xml:"EmailContacts"`
							Edition       string `xml:"Edition"`
							IdOwner       string `xml:"IdOwner"`
							LeagueType    string `xml:"LeagueType"`
							Logo          string `xml:"Logo"`
							ID            struct {
								Text  string `xml:",chardata"`
								Value string `xml:"Value"`
							} `xml:"Id"`
							MarketPlaceRookieGenerated string `xml:"MarketPlaceRookieGenerated"`
							HasPassword                string `xml:"HasPassword"`
							Solo                       string `xml:"Solo"`
							Description                string `xml:"Description"`
							MarketPlace                string `xml:"MarketPlace"`
							NeedValidation             string `xml:"NeedValidation"`
							Phase                      string `xml:"Phase"`
							Password                   string `xml:"Password"`
							Data                       string `xml:"Data"`
							Name                       string `xml:"Name"`
							Official                   string `xml:"Official"`
							Websites                   string `xml:"Websites"`
							TreasurySp                 string `xml:"TreasurySp"`
							NbRegisteredTeams          string `xml:"NbRegisteredTeams"`
							Flags                      string `xml:"Flags"`
						} `xml:"RowLeague"`
						Team struct {
							Text           string `xml:",chardata"`
							LevelupPending string `xml:"LevelupPending"`
							Statistics     struct {
								Text           string `xml:",chardata"`
								StatisticsTeam []struct {
									Text     string `xml:",chardata"`
									Category string `xml:"Category"`
									Row      struct {
										Text                   string `xml:",chardata"`
										IdLadder               string `xml:"IdLadder"`
										InflictedPasses        string `xml:"InflictedPasses"`
										Active                 string `xml:"Active"`
										InflictedCasualties    string `xml:"InflictedCasualties"`
										SustainedInterceptions string `xml:"SustainedInterceptions"`
										OccupationOwn          string `xml:"OccupationOwn"`
										AverageCashEarned      string `xml:"AverageCashEarned"`
										InflictedMetersPassing string `xml:"InflictedMetersPassing"`
										Category               string `xml:"Category"`
										Draws                  string `xml:"Draws"`
										SustainedCatches       string `xml:"SustainedCatches"`
										InflictedTackles       string `xml:"InflictedTackles"`
										SustainedTackles       string `xml:"SustainedTackles"`
										InflictedKO            string `xml:"InflictedKO"`
										IdLeague               string `xml:"IdLeague"`
										Points                 string `xml:"Points"`
										SustainedInjuries      string `xml:"SustainedInjuries"`
										InflictedDead          string `xml:"InflictedDead"`
										SustainedDead          string `xml:"SustainedDead"`
										MatchPlayed            string `xml:"MatchPlayed"`
										PossessionBall         string `xml:"PossessionBall"`
										CashEarned             string `xml:"CashEarned"`
										SustainedKO            string `xml:"SustainedKO"`
										MVP                    string `xml:"MVP"`
										AverageSpectators      string `xml:"AverageSpectators"`
										InflictedInterceptions string `xml:"InflictedInterceptions"`
										InflictedPushOuts      string `xml:"InflictedPushOuts"`
										LadderScore            string `xml:"LadderScore"`
										IdTeamListing          string `xml:"IdTeamListing"`
										SustainedMetersPassing string `xml:"SustainedMetersPassing"`
										InflictedCatches       string `xml:"InflictedCatches"`
										SustainedCasualties    string `xml:"SustainedCasualties"`
										SustainedPasses        string `xml:"SustainedPasses"`
										InflictedInjuries      string `xml:"InflictedInjuries"`
										ID                     string `xml:"ID"`
										Loss                   string `xml:"Loss"`
										SustainedTouchdowns    string `xml:"SustainedTouchdowns"`
										OccupationTheir        string `xml:"OccupationTheir"`
										SustainedMetersRunning string `xml:"SustainedMetersRunning"`
										Wins                   string `xml:"Wins"`
										LadderBestScore        string `xml:"LadderBestScore"`
										IdCompetition          string `xml:"IdCompetition"`
										InflictedMetersRunning string `xml:"InflictedMetersRunning"`
										InflictedTouchdowns    string `xml:"InflictedTouchdowns"`
									} `xml:"Row"`
								} `xml:"StatisticsTeam"`
							} `xml:"Statistics"`
							IdCoach   string `xml:"IdCoach"`
							TeamCards struct {
								Text            string `xml:",chardata"`
								TeamCardsByType []struct {
									Text      string `xml:",chardata"`
									TeamCards struct {
										Text     string `xml:",chardata"`
										TeamCard []struct {
											Text    string `xml:",chardata"`
											URL     string `xml:"Url"`
											RowCard struct {
												Text            string `xml:",chardata"`
												IdCardTypes     string `xml:"IdCardTypes"`
												StaticDataName  string `xml:"StaticDataName"`
												LocaDescription string `xml:"LocaDescription"`
												DataConstant    string `xml:"DataConstant"`
												LocaName        string `xml:"LocaName"`
												ID              string `xml:"ID"`
											} `xml:"RowCard"`
											IdTeam     string `xml:"IdTeam"`
											IdCard     string `xml:"IdCard"`
											IdTeamCard string `xml:"IdTeamCard"`
										} `xml:"TeamCard"`
									} `xml:"TeamCards"`
									MaxCardOfThisType string `xml:"MaxCardOfThisType"`
									IdCardType        string `xml:"IdCardType"`
								} `xml:"TeamCardsByType"`
							} `xml:"TeamCards"`
							MainRanking struct {
								Text          string `xml:",chardata"`
								Loss          string `xml:"Loss"`
								IdLadder      string `xml:"IdLadder"`
								Draws         string `xml:"Draws"`
								Score         string `xml:"Score"`
								Season        string `xml:"Season"`
								Rank          string `xml:"Rank"`
								IdLeague      string `xml:"IdLeague"`
								LeagueName    string `xml:"LeagueName"`
								LeaguePhase   string `xml:"LeaguePhase"`
								IdLeagueGroup string `xml:"IdLeagueGroup"`
								Wins          string `xml:"Wins"`
								LayoutType    string `xml:"LayoutType"`
								GroupIndex    string `xml:"GroupIndex"`
								IdTeam        struct {
									Text  string `xml:",chardata"`
									Value string `xml:"Value"`
								} `xml:"IdTeam"`
							} `xml:"MainRanking"`
							Rankings      string `xml:"Rankings"`
							NbPlayers     string `xml:"NbPlayers"`
							IdCompetition struct {
								Text  string `xml:",chardata"`
								Value string `xml:"Value"`
							} `xml:"IdCompetition"`
							IdLeague struct {
								Text  string `xml:",chardata"`
								Value string `xml:"Value"`
							} `xml:"IdLeague"`
							CompetitionRegistrationStatus string `xml:"CompetitionRegistrationStatus"`
							Online                        string `xml:"Online"`
							Validated                     string `xml:"Validated"`
							StadiumLevel                  string `xml:"StadiumLevel"`
							Row                           struct {
								Text                     string `xml:",chardata"`
								LevelupPending           string `xml:"LevelupPending"`
								StadiumName              string `xml:"StadiumName"`
								IdRaces                  int    `xml:"IdRaces"`
								Campaign                 string `xml:"Campaign"`
								Solo                     string `xml:"Solo"`
								Ai                       string `xml:"Ai"`
								StadiumInfrastructure    string `xml:"StadiumInfrastructure"`
								IdCheerleadersRace       string `xml:"IdCheerleadersRace"`
								IdOwner                  string `xml:"IdOwner"`
								Value                    string `xml:"Value"`
								Apothecary               string `xml:"Apothecary"`
								Logo                     string `xml:"Logo"`
								Validated                string `xml:"Validated"`
								StadiumLevel             string `xml:"StadiumLevel"`
								Predefined               string `xml:"Predefined"`
								Balms                    string `xml:"Balms"`
								Cheerleaders             string `xml:"Cheerleaders"`
								StadiumType              string `xml:"StadiumType"`
								Online                   string `xml:"Online"`
								IdOnlineSynchronizedFrom struct {
									Text  string `xml:",chardata"`
									Value string `xml:"Value"`
								} `xml:"IdOnlineSynchronizedFrom"`
								Edited     string `xml:"Edited"`
								NbPlayers  string `xml:"NbPlayers"`
								Deleted    string `xml:"Deleted"`
								Background string `xml:"Background"`
								Active     string `xml:"Active"`
								ID         struct {
									Text  string `xml:",chardata"`
									Value string `xml:"Value"`
								} `xml:"ID"`
								Leitmotiv                string `xml:"Leitmotiv"`
								IdCoach                  string `xml:"IdCoach"`
								Name                     string `xml:"Name"`
								AssistantCoaches         string `xml:"AssistantCoaches"`
								Popularity               string `xml:"Popularity"`
								Cash                     string `xml:"Cash"`
								Rerolls                  string `xml:"Rerolls"`
								Flags                    string `xml:"Flags"`
								PredefinedTeamCampaignId string `xml:"PredefinedTeamCampaignId"`
								TeamColor                string `xml:"TeamColor"`
							} `xml:"Row"`
						} `xml:"Team"`
						CoachName string `xml:"CoachName"`
					} `xml:"TeamRoster"`
				} `xml:"TeamRosters"`
				Server           string `xml:"Server"`
				Port             string `xml:"Port"`
				IsSSL            string `xml:"IsSSL"`
				LevelCabalVision string `xml:"LevelCabalVision"`
				StructStadium    string `xml:"StructStadium"`
				OpponentIdTeam   struct {
					Text  string `xml:",chardata"`
					Value string `xml:"Value"`
				} `xml:"OpponentIdTeam"`
				IdLeague struct {
					Text  string `xml:",chardata"`
					Value string `xml:"Value"`
				} `xml:"IdLeague"`
				RowTeamOpponent struct {
					Text                     string `xml:",chardata"`
					LevelupPending           string `xml:"LevelupPending"`
					StadiumName              string `xml:"StadiumName"`
					IdRaces                  int    `xml:"IdRaces"`
					Campaign                 string `xml:"Campaign"`
					Solo                     string `xml:"Solo"`
					Ai                       string `xml:"Ai"`
					StadiumInfrastructure    string `xml:"StadiumInfrastructure"`
					IdCheerleadersRace       string `xml:"IdCheerleadersRace"`
					IdOwner                  string `xml:"IdOwner"`
					Value                    string `xml:"Value"`
					Apothecary               string `xml:"Apothecary"`
					Logo                     string `xml:"Logo"`
					Validated                string `xml:"Validated"`
					StadiumLevel             string `xml:"StadiumLevel"`
					Predefined               string `xml:"Predefined"`
					Balms                    string `xml:"Balms"`
					Cheerleaders             string `xml:"Cheerleaders"`
					StadiumType              string `xml:"StadiumType"`
					Online                   string `xml:"Online"`
					IdOnlineSynchronizedFrom struct {
						Text  string `xml:",chardata"`
						Value string `xml:"Value"`
					} `xml:"IdOnlineSynchronizedFrom"`
					Edited     string `xml:"Edited"`
					NbPlayers  string `xml:"NbPlayers"`
					Deleted    string `xml:"Deleted"`
					Background string `xml:"Background"`
					Active     string `xml:"Active"`
					ID         struct {
						Text  string `xml:",chardata"`
						Value string `xml:"Value"`
					} `xml:"ID"`
					Leitmotiv                string `xml:"Leitmotiv"`
					IdCoach                  string `xml:"IdCoach"`
					Name                     string `xml:"Name"`
					AssistantCoaches         string `xml:"AssistantCoaches"`
					Popularity               string `xml:"Popularity"`
					Cash                     string `xml:"Cash"`
					Rerolls                  string `xml:"Rerolls"`
					Flags                    string `xml:"Flags"`
					PredefinedTeamCampaignId string `xml:"PredefinedTeamCampaignId"`
					TeamColor                string `xml:"TeamColor"`
				} `xml:"RowTeamOpponent"`
				IdCoachHome                string `xml:"IdCoachHome"`
				Stadium                    string `xml:"Stadium"`
				CoachHomeOwnedDLC          string `xml:"CoachHomeOwnedDLC"`
				IsSSLWebsockets            string `xml:"IsSSLWebsockets"`
				DataGameInvitationSpecific string `xml:"DataGameInvitationSpecific"`
				DrawAllowed                string `xml:"DrawAllowed"`
				IdGame                     string `xml:"IdGame"`
				RowTeamSelf                struct {
					Text                     string `xml:",chardata"`
					LevelupPending           string `xml:"LevelupPending"`
					StadiumName              string `xml:"StadiumName"`
					IdRaces                  int    `xml:"IdRaces"`
					Campaign                 string `xml:"Campaign"`
					Solo                     string `xml:"Solo"`
					Ai                       string `xml:"Ai"`
					StadiumInfrastructure    string `xml:"StadiumInfrastructure"`
					IdCheerleadersRace       string `xml:"IdCheerleadersRace"`
					IdOwner                  string `xml:"IdOwner"`
					Value                    string `xml:"Value"`
					Apothecary               string `xml:"Apothecary"`
					Logo                     string `xml:"Logo"`
					Validated                string `xml:"Validated"`
					StadiumLevel             string `xml:"StadiumLevel"`
					Predefined               string `xml:"Predefined"`
					Balms                    string `xml:"Balms"`
					Cheerleaders             string `xml:"Cheerleaders"`
					StadiumType              string `xml:"StadiumType"`
					Online                   string `xml:"Online"`
					IdOnlineSynchronizedFrom struct {
						Text  string `xml:",chardata"`
						Value string `xml:"Value"`
					} `xml:"IdOnlineSynchronizedFrom"`
					Edited     string `xml:"Edited"`
					NbPlayers  string `xml:"NbPlayers"`
					Deleted    string `xml:"Deleted"`
					Background string `xml:"Background"`
					Active     string `xml:"Active"`
					ID         struct {
						Text  string `xml:",chardata"`
						Value string `xml:"Value"`
					} `xml:"ID"`
					Leitmotiv                string `xml:"Leitmotiv"`
					IdCoach                  string `xml:"IdCoach"`
					Name                     string `xml:"Name"`
					AssistantCoaches         string `xml:"AssistantCoaches"`
					Popularity               string `xml:"Popularity"`
					Cash                     string `xml:"Cash"`
					Rerolls                  string `xml:"Rerolls"`
					Flags                    string `xml:"Flags"`
					PredefinedTeamCampaignId string `xml:"PredefinedTeamCampaignId"`
					TeamColor                string `xml:"TeamColor"`
				} `xml:"RowTeamSelf"`
				Started string `xml:"Started"`
				IdMatch struct {
					Text  string `xml:",chardata"`
					Value string `xml:"Value"`
				} `xml:"IdMatch"`
				Ranked         string `xml:"Ranked"`
				IdCoachAway    string `xml:"IdCoachAway"`
				PortWebsockets string `xml:"PortWebsockets"`
				NameStadium    string `xml:"NameStadium"`
				IdTeamHome     struct {
					Text  string `xml:",chardata"`
					Value string `xml:"Value"`
				} `xml:"IdTeamHome"`
				CoachAwayOwnedDLC string `xml:"CoachAwayOwnedDLC"`
				CoachHomeReady    string `xml:"CoachHomeReady"`
				IdSession         string `xml:"IdSession"`
				IdCompetition     struct {
					Text  string `xml:",chardata"`
					Value string `xml:"Value"`
				} `xml:"IdCompetition"`
				LevelStadium string `xml:"LevelStadium"`
				RowLeague    struct {
					Text          string `xml:",chardata"`
					EmailContacts string `xml:"EmailContacts"`
					Edition       string `xml:"Edition"`
					IdOwner       string `xml:"IdOwner"`
					LeagueType    string `xml:"LeagueType"`
					Logo          string `xml:"Logo"`
					ID            struct {
						Text  string `xml:",chardata"`
						Value string `xml:"Value"`
					} `xml:"Id"`
					MarketPlaceRookieGenerated string `xml:"MarketPlaceRookieGenerated"`
					HasPassword                string `xml:"HasPassword"`
					Solo                       string `xml:"Solo"`
					Description                string `xml:"Description"`
					MarketPlace                string `xml:"MarketPlace"`
					NeedValidation             string `xml:"NeedValidation"`
					Phase                      string `xml:"Phase"`
					Password                   string `xml:"Password"`
					Data                       string `xml:"Data"`
					Name                       string `xml:"Name"`
					Official                   string `xml:"Official"`
					Websites                   string `xml:"Websites"`
					TreasurySp                 string `xml:"TreasurySp"`
					NbRegisteredTeams          string `xml:"NbRegisteredTeams"`
					Flags                      string `xml:"Flags"`
				} `xml:"RowLeague"`
				TurnDuration string `xml:"TurnDuration"`
				IdTeamAway   struct {
					Text  string `xml:",chardata"`
					Value string `xml:"Value"`
				} `xml:"IdTeamAway"`
				IdTeam struct {
					Text  string `xml:",chardata"`
					Value string `xml:"Value"`
				} `xml:"IdTeam"`
				CoachAwayReady string `xml:"CoachAwayReady"`
				OpponentId     string `xml:"OpponentId"`
				RowCompetition struct {
					Text              string `xml:",chardata"`
					IdPreviousEdition struct {
						Text  string `xml:",chardata"`
						Value string `xml:"Value"`
					} `xml:"IdPreviousEdition"`
					IdOwner                string `xml:"IdOwner"`
					RankedChallengeAllowed string `xml:"RankedChallengeAllowed"`
					Logo                   string `xml:"Logo"`
					RankedMatchAllowed     string `xml:"RankedMatchAllowed"`
					CheckinPrice           string `xml:"CheckinPrice"`
					ID                     struct {
						Text  string `xml:",chardata"`
						Value string `xml:"Value"`
					} `xml:"Id"`
					AutomatchAllowed         string `xml:"AutomatchAllowed"`
					KeepReplayPolicy         string `xml:"KeepReplayPolicy"`
					CloneOnNextLeagueEdition string `xml:"CloneOnNextLeagueEdition"`
					LeaguePhase              string `xml:"LeaguePhase"`
					ChallengeAllowed         string `xml:"ChallengeAllowed"`
					IdLeague                 struct {
						Text  string `xml:",chardata"`
						Value string `xml:"Value"`
					} `xml:"IdLeague"`
					NbRegisteredTeams        string `xml:"NbRegisteredTeams"`
					AllowSelfMatchValidation string `xml:"AllowSelfMatchValidation"`
					PrematchEvents           string `xml:"PrematchEvents"`
					Flags                    string `xml:"Flags"`
					Rewards                  string `xml:"Rewards"`
					Description              string `xml:"Description"`
					IdCompetitionTypes       string `xml:"IdCompetitionTypes"`
					RegistrationMode         string `xml:"RegistrationMode"`
					NbRounds                 string `xml:"NbRounds"`
					NbTeamsMax               string `xml:"NbTeamsMax"`
					CompetitionStatus        string `xml:"CompetitionStatus"`
					AdmissionStadiumLevel    string `xml:"AdmissionStadiumLevel"`
					UnRankedChallengeAllowed string `xml:"UnRankedChallengeAllowed"`
					Name                     string `xml:"Name"`
					CurrentRound             string `xml:"CurrentRound"`
					LeagueEdition            string `xml:"LeagueEdition"`
					UnRankedMatchAllowed     string `xml:"UnRankedMatchAllowed"`
					PureAutomatch            string `xml:"PureAutomatch"`
					NeedValidation           string `xml:"NeedValidation"`
					AcceptTicketRequest      string `xml:"AcceptTicketRequest"`
					TurnDuration             string `xml:"TurnDuration"`
					GameOptions              string `xml:"GameOptions"`
				} `xml:"RowCompetition"`
			} `xml:"BBGameSession"`
		} `xml:"MessageData"`
	} `xml:"Specific"`
	Host         string `xml:"Host"`
	OwnedDLCHome string `xml:"OwnedDLCHome"`
}
