// Conic Gradient Types

export interface ConicStop {
	/** The legend label value. */
	label?: string;
	/** The desired color value. */
	color: string;
	/** Starting stop value (0-100) */
	start: number;
	/** Ending stop value (0-100) */
	end: number;
}
