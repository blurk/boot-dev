def unlock_achievement(before_xp, ach_xp, ach_name):
		after_xp = before_xp + ach_xp
		text = f"Achievement Unlocked: {ach_name}"
		return after_xp, text